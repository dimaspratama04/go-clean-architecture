package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

func main() {
	godotenv.Load()
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Username: os.Getenv("REDIS_USER"),
		Password: os.Getenv("REDIS_PASSWORD"),
		// TLSConfig: &tls.Config{},
		DB: 0,
	})

	if err := rdb.Ping(ctx).Err(); err != nil {
		log.Fatalf("failed to connect redis: %v", err)
	}

	log.Println("✅ Connected to redis")

	stream := "transaction"
	group := "trx-workers"
	consumer := "worker-1"

	// Create group if not exists
	_ = rdb.XGroupCreateMkStream(ctx, stream, group, "0").Err()

	for {
		// Read messages
		resp, err := rdb.XReadGroup(ctx, &redis.XReadGroupArgs{
			Group:    group,
			Consumer: consumer,
			Streams:  []string{stream, ">"},
			Count:    10,
			Block:    5 * time.Second,
		}).Result()

		if err != nil && err != redis.Nil {
			log.Println("read err:", err)
			continue
		}

		if len(resp) == 0 {
			continue
		}

		for _, msg := range resp[0].Messages {
			fmt.Printf("[Worker] Processing msg ID=%s values=%v\n", msg.ID, msg.Values)

			// Simulate processing
			time.Sleep(500 * time.Millisecond)

			// ACK
			if err := rdb.XAck(ctx, stream, group, msg.ID).Err(); err != nil {
				log.Println("ack err:", err)
			}
		}
	}
}
