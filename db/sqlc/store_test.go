package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransTx(t *testing.T) {
	store := NewStore(testDB)

	account1 := createAccountRandom(t)
	account2 := createAccountRandom(t)

	//The best way to make sure that our transaction works is to run it concurrently
	//We need to run it with several concurrent goroutines.
	//Run n concurrent transfer transactions.
	n := 5
	amount := int64(10)

	// The better way to verify the result or an error is to use channel
	errsChan := make(chan error)
	resultsChan := make(chan TransferTxResult)

	// Creating transactions
	for i := 0; i < n; i++ {
		go func() {
			result, err := store.TransferTx(context.Background(), TransferTxParams{
				FromAccountID: account1.ID,
				ToAccountID:   account2.ID,
				Amount:        amount,
			})
			errsChan <- err
			resultsChan <- result
		}()
	}
	//Testing the result
	for i := 0; i < n; i++ {
		err := <-errsChan
		require.NoError(t, err)

		result := <-resultsChan
		require.NotEmpty(t, result)
		//Check Transfer
		transfer := result.Transfer
		require.NotEmpty(t, transfer)
		require.Equal(t, account1.ID, transfer.FromAccountID)
		require.Equal(t, account2.ID, transfer.ToAccountID)
		require.Equal(t, amount, transfer.Amount)
		require.NotZero(t, transfer.ID)

		_, err = store.GetTransfer(context.Background(), transfer.ID)
		require.NoError(t, err)

		//Check FromEntries
		fromEntry := result.FromEntry
		require.NotEmpty(t, fromEntry)
		require.Equal(t, account1.ID, fromEntry.AccountID)
		require.Equal(t, -amount, fromEntry.Amount)
		require.NotZero(t, fromEntry.ID)

		//Check ToEntries
		toEntry := result.ToEntry
		require.NotEmpty(t, toEntry)
		require.Equal(t, account2.ID, toEntry.AccountID)
		require.Equal(t, amount, toEntry.Amount)
		require.NotZero(t, toEntry.ID)
	}

}
