package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Store struct {
	// We need to provide Queries struct. However, each query only do ONE operation on ONE specific table.
	// So Queries doesn't support transaction. Transaction needs to ALTER MULTIPLE tables.
	// This is called composition
	*Queries
	db *pgxpool.Pool
}

// NewStore is a constructor function for creating a new Store.
// It takes a pointer to a pgxpool.Pool as an argument, which is a connection pool for a PostgreSQL database.
func NewStore(db *pgxpool.Pool) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

// ExecTx executes a function within a database transaction.
func (store *Store) excTx(ctx context.Context, fn func(*Queries) error) error {
	//Starts a transaction
	tx, err := store.db.Begin(ctx)
	if err != nil {
		return err
	}
	//Execute query
	q := New(tx)
	//Check if there is any error with that query
	err = fn(q)
	// Inside the function, we can apply RollBack
	if err != nil {
		if rBaErr := tx.Rollback(ctx); rBaErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rBaErr)
		}
	}

	return tx.Commit(ctx)
}

type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}
type TransferTxResult struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
}

// Performs a money transfer from one account to another.
// It create a
// 1- Transfer record,
// 2- add account entries,
// 3- update account's balance within a single database transaction.
func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult

	err := store.excTx(ctx, func(q *Queries) error {
		var err error
		//1- Create a transfer record
		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: arg.FromAccountID,
			ToAccountID:   arg.ToAccountID,
			Amount:        arg.Amount,
		})
		if err != nil {
			return err
		}
		//2- Create FROM entry
		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.FromAccountID,
			Amount:    -arg.Amount,
		})
		if err != nil {
			return err
		}
		//2- Create TO entry
		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.ToAccountID,
			Amount:    arg.Amount,
		})
		if err != nil {
			return err
		}
		//Update accounts It involves locking and preventing potential deadlocks

		//3- Update FROM account's balance
		account1, err := q.GetAccount(ctx, arg.FromAccountID)
		if err != nil {
			return err
		}
		err = q.UpdateAccount(ctx, UpdateAccountParams{
			ID:      arg.FromAccountID,
			Balance: account1.Balance - arg.Amount,
		})
		if err != nil {
			return err
		}
		account2, err := q.GetAccount(ctx, arg.ToAccountID)
		if err != nil {
			return err
		}
		err = q.UpdateAccount(ctx, UpdateAccountParams{
			ID:      arg.ToAccountID,
			Balance: account2.Balance + arg.Amount,
		})
		if err != nil {
			return err
		}
		result.FromAccount, err = q.GetAccount(ctx, arg.FromAccountID)
		if err != nil {
			return err
		}
		result.ToAccount, err = q.GetAccount(ctx, arg.ToAccountID)
		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}
