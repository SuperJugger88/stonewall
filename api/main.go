package main

import (
	"api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", routes.MainPage)

	router.Run(":80")
}
