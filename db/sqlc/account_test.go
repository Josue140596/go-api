package db

import (
	"context"
	"go/api/utils"
	"testing"

	"github.com/stretchr/testify/require"
)

/**
* TESTS
** To test all CRUD operations, we need to create a new account,
** but we should make sure that they are independent of each other.
** Why? Because it would be very hard to maintain.
** For this reason , each test should create its own account records.
**/

func createAccountRandom(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    utils.RandomOwner(),
		Balance:  utils.RandomMoney(),
		Currency: utils.RandomCurrency(),
	}
	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)
	//Match arg with new created account
	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	//ID must be greater than 0
	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
	return account
}

func TestCreateAccount(t *testing.T) {
	createAccountRandom(t)
}

func TestGetAccount(t *testing.T) {
	accountCreated := createAccountRandom(t)
	account, error := testQueries.GetAccount(context.Background(), accountCreated.ID)
	require.NoError(t, error)
	require.NotEmpty(t, account)

	require.Equal(t, accountCreated.ID, account.ID)
	require.Equal(t, accountCreated.Owner, account.Owner)
	require.Equal(t, accountCreated.Balance, account.Balance)
	require.Equal(t, accountCreated.Currency, account.Currency)
}

func TestUpdateAccount(t *testing.T) {
	accountCreated := createAccountRandom(t)
	arg := UpdateAccountParams{
		ID:      accountCreated.ID,
		Balance: utils.RandomMoney(),
	}
	error := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, error)
	require.NotEqual(t, accountCreated.Balance, arg.Balance)
}

func TestDeleteAccount(t *testing.T) {
	accountCreated := createAccountRandom(t)
	error := testQueries.DeleteAccount(context.Background(), accountCreated.ID)
	require.NoError(t, error)
	//Check if account is deleted
	accountDeleted, error := testQueries.GetAccount(context.Background(), accountCreated.ID)
	require.Error(t, error)
	require.Empty(t, accountDeleted)
}

func TestListAccount(t *testing.T) {
	for i := 0; i < 10; i++ {
		createAccountRandom(t)
	}
	arg := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}
	accounts, error := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, error)
	require.Len(t, accounts, 5)

	for _, v := range accounts {
		require.NotEmpty(t, v)
	}
}
