package redis

import (
	"context"
	"time"
)

type Redis interface {
	SetEX(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	Get(ctx context.Context, key string, value interface{}) error
	Delete(ctx context.Context, key string) error
}
