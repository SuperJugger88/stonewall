package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func handler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Logema Squad",
	})
}

func main() {
	router := gin.Default()

	router.GET("/", handler)

	err := router.Run(":80")
	if err != nil {
		log.Fatal(err)
	}
}
