package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"gitongaraymond/bank_backend_sample/util"
	"testing"
	"time"
)

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}
func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	account1 := createRandomAccount(t)

	account2, err := testQueries.GetAccount(context.Background(), account1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Currency, account2.Currency)
	require.Equal(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.Owner, account2.Owner)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	account1 := createRandomAccount(t)

	arg := UpdateAccountParams{account1.ID, util.RandomMoney()}
	err := testQueries.UpdateAccount(context.Background(), arg)

	account2, err1 := testQueries.GetAccount(context.Background(), account1.ID)

	require.NoError(t, err)
	require.NoError(t, err1)
	require.Equal(t, arg.Balance, account2.Balance)
	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Currency, account2.Currency)
	require.Equal(t, account1.Owner, account2.Owner)

}

func TestDeleteAccount(t *testing.T) {
	account1 := createRandomAccount(t)

	err := testQueries.DeleteAccount(context.Background(), account1.ID)

	account2, _ := testQueries.GetAccount(context.Background(), account1.ID)

	require.NoError(t, err)
	require.Empty(t, account2.Balance)
	require.Empty(t, account2.ID)
	require.Empty(t, account2.Currency)
	require.Empty(t, account2.Owner)

}

func TestListAccounts(t *testing.T) {
	createRandomAccount(t)
	createRandomAccount(t)

	args := ListAccountsParams{2, 2}
	accounts, err := testQueries.ListAccounts(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, accounts)
	require.Equal(t, len(accounts), 2)

}
