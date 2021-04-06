package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/currenttime", currentTime)

	router.Run(":8081")
}

func currentTime(c *gin.Context) {
	currentTime := time.Now()
	c.JSON(http.StatusOK, gin.H{
		"time": currentTime.Format("2006-01-02 3:4:5 pm"),
	},
	)
}
