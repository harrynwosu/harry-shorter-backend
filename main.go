package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Start the web server
func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "HarryShortner Go URL Shortner!",
		})
	})

	err := router.Run(":9808")
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}
}