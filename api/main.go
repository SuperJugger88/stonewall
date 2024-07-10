package main

import (
	"api/routes"
	"context"
	crdbpgx "github.com/cockroachdb/cockroach-go/v2/crdb/crdbpgxv5"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"log"
	"os"
)

func initRouter() {
	router := gin.Default()

	router.GET("/api/v1/test", routes.GetMainPage)

	err := router.Run(os.Getenv("APP_URL"))
	if err != nil {
		panic(err)
	}
}

func initTable(ctx context.Context, tx pgx.Tx) error {
	log.Println("Drop existing accounts table if necessary.")
	if _, err := tx.Exec(ctx, "DROP TABLE IF EXISTS accounts"); err != nil {
		return err
	}

	// Create the accounts table
	log.Println("Creating accounts table.")
	if _, err := tx.Exec(ctx,
		"CREATE TABLE accounts (id UUID PRIMARY KEY DEFAULT gen_random_uuid(), balance INT8)"); err != nil {
		return err
	}
	return nil
}

func insertRows(ctx context.Context, tx pgx.Tx, accts [4]uuid.UUID) error {
	// Insert four rows into the "accounts" table.
	log.Println("Creating new rows...")
	if _, err := tx.Exec(ctx,
		"INSERT INTO accounts (id, balance) VALUES ($1, $2), ($3, $4), ($5, $6), ($7, $8)", accts[0], 250, accts[1], 100, accts[2], 500, accts[3], 300); err != nil {
		return err
	}
	return nil
}

func initDatabaseConnection() {
	config, err := pgx.ParseConfig(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	config.RuntimeParams["application_name"] = "$ docs_simplecrud_gopgx"
	conn, err := pgx.ConnectConfig(context.Background(), config)
	if err != nil {
		panic(err)
	}
	defer conn.Close(context.Background())

	err = crdbpgx.ExecuteTx(context.Background(), conn, pgx.TxOptions{}, func(tx pgx.Tx) error {
		return initTable(context.Background(), tx)
	})

	var accounts [4]uuid.UUID
	for i := 0; i < len(accounts); i++ {
		accounts[i] = uuid.New()
	}

	err = crdbpgx.ExecuteTx(context.Background(), conn, pgx.TxOptions{}, func(tx pgx.Tx) error {
		return insertRows(context.Background(), tx, accounts)
	})
	if err == nil {
		log.Println("New rows created.")
	} else {
		log.Fatal("error: ", err)
	}
}

func main() {
	initRouter()
	initDatabaseConnection()
}
