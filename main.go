package main

import (
	"github.com/gin-gonic/gin"
	"mock-interview-back/app/handlers"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/authorization", handlers.Authorization)

	err := r.Run()
	if err != nil {
		return
	}
}
