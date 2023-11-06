package db

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

var testQueries *Queries

const dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"

func TestMain(m *testing.M) {
	conn, err := pgxpool.New(context.Background(), dbSource)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	testQueries = New(conn)

	os.Exit(m.Run())
}
