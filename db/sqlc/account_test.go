package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/hanginfty/simple-bank/util"
	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {
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
}

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	return account
}

func TestGetAccount(t *testing.T) {
	account := createRandomAccount(t)

	got, err := testQueries.GetAccount(context.Background(), account.ID)
	require.NoError(t, err)
	require.Equal(t, account.ID, got.ID)
	require.Equal(t, account.Balance, got.Balance)
	require.Equal(t, account.Currency, got.Currency)
	require.WithinDuration(t, account.CreatedAt, got.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	account := createRandomAccount(t)

	// mutate the account.Blance

	arg := UpdateAccountParams{
		ID:      account.ID,
		Balance: util.RandomMoney(),
	}

	got, err := testQueries.UpdateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, got)

	require.Equal(t, account.ID, got.ID)
	require.Equal(t, account.Owner, got.Owner)
	require.Equal(t, arg.Balance, got.Balance)
	require.Equal(t, account.Currency, got.Currency)
	require.WithinDuration(t, account.CreatedAt, got.CreatedAt, time.Second)

}

func TestDeleteAccount(t *testing.T) {
	account := createRandomAccount(t)

	err := testQueries.DeleteAccount(context.Background(), account.ID)
	require.NoError(t, err)

	got, err := testQueries.GetAccount(context.Background(), account.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, got)
}

func TestListAccounts(t *testing.T) {
	accounts := []Account{
		createRandomAccount(t),
		createRandomAccount(t),
		createRandomAccount(t),
	}

	arg := ListAccountsParams{
		Limit:  3,
		Offset: 43,
	}

	got, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, accounts, got)
	require.Len(t, got, 3)

	for _, got := range got {
		require.NotEmpty(t, got)
	}
}
