package db

import "testing"

func TestTransTx(t *testing.T) {
	store := NewStore(testDB)

	account1 := createAccountRandom(t)
	account2 := createAccountRandom(t)

	//Run n concurrent transfer transactions
	n := 5
	amount := int64(10)

}
