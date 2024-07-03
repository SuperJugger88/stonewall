package main

import (
	"api/routes"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"os"
)

func initRouter() {
	router := gin.Default()

	router.GET("/", routes.GetMainPage)

	err := router.Run(os.Getenv("APP_URL"))
	if err != nil {
		panic(err)
	}
}

func initDatabaseConnection() {
	config, err := pgx.ParseConfig(os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	conn, err := pgx.ConnectConfig(context.Background(), config)
	if err != nil {
		panic(err)
	}
	defer conn.Close(context.Background())
}

func main() {
	initRouter()
	initDatabaseConnection()
}
