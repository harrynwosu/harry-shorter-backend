package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/harrynwosu/harryshorter/redis_store"
	"github.com/harrynwosu/harryshorter/shortener"
)

type GenerateShortUrlRequest struct {
	// binding: required => mandatory/required fields
	OriginalUrl string `json:"original_url" binding:"required"`
	UserId string `json:"user_id" binding:"required"`
}

func GenerateShortUrl(c *gin.Context) {
	var request GenerateShortUrlRequest
	// Error handling for when OriginalUrl or UserId is missing from request
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shortUrl := shortener.GenerateShortLink(request.OriginalUrl, request.UserId)
	redis_store.SaveUrlMapping(shortUrl, request.OriginalUrl, request.UserId)

	host := "http://localhost:9808/"
	c.JSON(200, gin.H{
		"message":   "short url created successfully",
		"short_url": host + shortUrl,
	})
}

func HandleUrlRedirect(c *gin.Context) {
	// Get shortURL from the URL params `/:shortUrl`
	// then redirect to the original gotten from the Redis store
	shortUrl := c.Param("shortUrl")
	originalUrl := redis_store.GetOriginalUrl(shortUrl)

	c.Redirect(302, originalUrl)
}