package db

import (
	"context"
	"fmt"
	"go/api/utils"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

var testQueries *Queries
var testDB *pgxpool.Pool

func TestMain(m *testing.M) {
	//Configuration
	conf, errorConfig := utils.LoadConfig("../..")
	if errorConfig != nil {
		panic(fmt.Errorf("fatal error config file: %w", errorConfig))
	}
	var err error
	testDB, err = pgxpool.New(context.Background(), conf.DB_SOURCE)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer testDB.Close()

	testQueries = New(testDB)

	os.Exit(m.Run())
}
