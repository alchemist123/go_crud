// main.go
package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a Gin router
	router := gin.Default()

	// Define a route handler
	router.GET("/", func(c *gin.Context) {
		fmt.Print("this a body of request", c.Request.Body)
		c.JSON(http.StatusOK, gin.H{"message": "Hello, You Created a Web App!"})
	})
	router.POST("/hello", func(ctx *gin.Context) {
		ctx.JSON(http.StatusAccepted, gin.H{"message": "hello"})
	})
	// Run the server on port 8080
	router.Run(":8080")
}
