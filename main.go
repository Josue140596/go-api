package main

import (
	"context"
	"fmt"
	"go/api/api"
	db "go/api/db/sqlc"
	"go/api/utils"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var conn *pgxpool.Pool

func main() {
	//Configuration
	conf, errorConfig := utils.LoadConfig(".")
	if errorConfig != nil {
		panic(fmt.Errorf("fatal error config file: %w", errorConfig))
	}

	var err error
	conn, err = pgxpool.New(context.Background(), conf.DB_SOURCE)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(conf.SERVER_ADDRESS)
	if err != nil {
		log.Fatal("Cannot start server: ", err)
	}
}
