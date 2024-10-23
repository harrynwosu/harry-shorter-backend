package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/harrynwosu/harryshorter/handler"
	"github.com/harrynwosu/harryshorter/redis_store"
)

// Start the web server
func main() {
	// Initialize our Redis store
	redis_store.InitializeStore()

	router := gin.Default()

	// Route handlers 🚀
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "HarryShortner Go URL Shortner!",
		})
	})

	router.POST("/generate-short-url", func(c *gin.Context) {
		handler.GenerateShortUrl(c)
	})

	router.GET("/:shortUrl", func(c *gin.Context) {
		handler.HandleUrlRedirect(c)
	})

	err := router.Run(":9808")
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}
}