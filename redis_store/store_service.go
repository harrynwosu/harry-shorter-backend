package redis_store

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type StorageService struct {
	redisClient *redis.Client
}

var (
	storeService = &StorageService{}
	ctx = context.Background()
)

// TODO: the cache duration shouldn't have an expiration time
// An LRU policy config should be set where the 
// values that are retrieved less often are purged automatically from 
// the cache and stored back in RDBMS whenever the cache is full
const CacheDuration = 6 * time.Hour

func InitializeStore () *StorageService {
	// Initialize new Redis client with default options
	newRedisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})

	// Ping the Redis server to verify the connection works
	pong, err := newRedisClient.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Error initializing Redis: %v", err))
	}

	fmt.Printf("\n Redis started successfully: pong message = {%s}", pong)
	storeService.redisClient = newRedisClient
	return storeService
}

// Storage API

// Save mapping of short url to original url
func SaveUrlMapping(shortUrl string, originalUrl string, userId string) {
	err := storeService.redisClient.Set(ctx, shortUrl, originalUrl, CacheDuration).Err()
	if err != nil {
		panic(fmt.Sprintf("Failed saving key url | Error: %v - shortUrl: %s - originalUrl: %s\n", err, shortUrl, originalUrl))
	}
}


// Retrieve original url given the short encoded url for appropriate redirects
func GetOriginalUrl(shortUrl string) string {
	result, err := storeService.redisClient.Get(ctx, shortUrl).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed RetrieveInitialUrl url | Error: %v - shortUrl: %s\n", err, shortUrl))
	}
	return result
}