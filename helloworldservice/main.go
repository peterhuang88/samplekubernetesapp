package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/helloworld", helloWorld)

	router.Run(":8082")
}

func helloWorld(c *gin.Context) {
	// currentTime := time.Now()
	c.JSON(http.StatusOK, gin.H{
		"message": "hello world",
	},
	)
}
