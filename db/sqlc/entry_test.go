package db

import (
	"context"
	"go/api/utils"
	"testing"

	"github.com/stretchr/testify/require"
)

func testCreateEntryRandom(t *testing.T, account Account) Entry {
	var arg = CreateEntryParams{
		AccountID: account.ID,
		Amount:    utils.RandomMoney(),
	}
	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, entry.AccountID, account.ID)
	require.Equal(t, arg.Amount, entry.Amount)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)
	return entry
}

func TestCreateEntry(t *testing.T) {
	account := createAccountRandom(t)
	testCreateEntryRandom(t, account)
}

func TestGetEntry(t *testing.T) {
	account := createAccountRandom(t)
	entryCreated := testCreateEntryRandom(t, account)
	entry, err := testQueries.GetEntry(context.Background(), entryCreated.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, entryCreated.ID, entry.ID)
	require.Equal(t, entryCreated.AccountID, entry.AccountID)
	require.Equal(t, entryCreated.Amount, entry.Amount)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)
}

func TestListEntries(t *testing.T) {
	account := createAccountRandom(t)
	for i := 0; i < 10; i++ {
		testCreateEntryRandom(t, account)
	}
	arg := ListEntriesParams{
		AccountID: account.ID,
		Limit:     5,
		Offset:    3,
	}
	entries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries, 5)
}
