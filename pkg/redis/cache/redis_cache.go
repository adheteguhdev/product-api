package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

func (c *RedisCache) SetEX(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	bytes, err := json.Marshal(value)
	if err != nil {
		msg := fmt.Sprintf("failed to marshal value: %v", err)
		c.logger.Errorf(msg)
		return fmt.Errorf(msg)
	}

	err = c.client.Set(ctx, key, bytes, expiration).Err()
	if err != nil {
		msg := fmt.Sprintf("failed to set key-value in Redis: %v", err)
		c.logger.Errorf(msg)
		return fmt.Errorf(msg)
	}
	c.logger.Infof("successfully set key-value in Redis")
	return nil
}

func (c *RedisCache) Get(ctx context.Context, key string, value interface{}) error {
	bytes, err := c.client.Get(ctx, key).Bytes()
	if err == redis.Nil {
		msg := fmt.Sprintf("key %s not found in Redis", key)
		c.logger.Errorf(msg)
		return fmt.Errorf(msg)
	} else if err != nil {
		msg := fmt.Sprintf("failed to get key-value from Redis: %v", err)
		c.logger.Errorf(msg)
		return fmt.Errorf(msg)
	}

	err = json.Unmarshal(bytes, value)
	if err != nil {
		msg := fmt.Sprintf("failed to unmarshal value: %v", err)
		c.logger.Errorf(msg)
		return fmt.Errorf(msg)
	}
	c.logger.Infof("key %s found in Redis", key)
	return nil
}

func (c *RedisCache) Delete(ctx context.Context, key string) error {
	err := c.client.Del(ctx, key).Err()
	if err != nil {
		msg := fmt.Sprintf("failed to delete key-value from Redis: %v", err)
		c.logger.Errorf(msg)
		return fmt.Errorf(msg)
	}
	c.logger.Infof("successfully delete key-value in Redis")
	return nil
}
