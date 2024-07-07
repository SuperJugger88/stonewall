package main

import (
	"api/routes"
	"context"
	crdbpgx "github.com/cockroachdb/cockroach-go/v2/crdb/crdbpgxv5"
	"github.com/gin-gonic/gin"
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

func initDatabaseConnection() {
	config, err := pgx.ParseConfig(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	conn, err := pgx.ConnectConfig(context.Background(), config)
	if err != nil {
		panic(err)
	}
	defer conn.Close(context.Background())

	err = crdbpgx.ExecuteTx(context.Background(), conn, pgx.TxOptions{}, func(tx pgx.Tx) error {
		return initTable(context.Background(), tx)
	})
}

func main() {
	initRouter()
	initDatabaseConnection()
}
