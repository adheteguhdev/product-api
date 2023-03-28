package cache

import (
	"product-api/pkg/logger"

	"github.com/go-redis/redis/v8"
)

type RedisCache struct {
	client *redis.Client
	logger *logger.Logger
}

func NewRedisCache(client *redis.Client, logger *logger.Logger) *RedisCache {

	return &RedisCache{
		client: client,
		logger: logger,
	}
}
