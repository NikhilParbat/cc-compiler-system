package main

import (
	"fmt"
	"net/http"

	"github.com/NikhilParbat/CC-Compiler-Go/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Use(func(c *gin.Context) {
		// Set CORS headers
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type")

		// Handle preflight options request
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		c.Next()
	})

	router.POST("/execute", func(c *gin.Context) {
		// Your actual request handling logic goes here
		controllers.ExecuteCode(c)
	})

	fmt.Println("Server listening on port 5000...")
	router.Run(":5000")
}
