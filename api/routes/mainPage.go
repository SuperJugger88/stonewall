package routes

import "github.com/gin-gonic/gin"

func MainPage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World",
	})
}
