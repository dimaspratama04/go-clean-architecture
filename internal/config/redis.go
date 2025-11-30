package config

import (
	"context"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

func NewRedisClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Username: os.Getenv("REDIS_USER"),
		Password: os.Getenv("REDIS_PASSWORD"),
		// TLSConfig: &tls.Config{},
		DB: 0})

	ctx := context.Background()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Printf("❌ Failed to connect to Redis: %v", err)
		log.Println("Continuing without Redis connection...")
		return nil
	}

	log.Println("✅ Connected to redis")
	return rdb
}
