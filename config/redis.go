package config

import (
	"context"
	"fmt"
	"os"
	"product-api/pkg/logger"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
)

var RedisClient *redis.Client

func InitRedis() error {
	log := logger.NewLogger()

	err := godotenv.Load()
	if err != nil {
		log.Errorf("failed to load .env file")
	}

	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")
	password := os.Getenv("REDIS_PASSWORD")

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: password,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pong, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		msg := fmt.Sprintf("failed to connect to Redis: %v", err)
		log.Errorf(msg)
		return fmt.Errorf(msg)

	}

	log.Infof("Redis connected: %v", pong)

	return nil
}
