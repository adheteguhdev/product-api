package order

import (
	"product-api/pkg/logger"
	"product-api/pkg/redis/cache"
	r "product-api/repository"

	"github.com/go-redis/redis/v8"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
)

type orderHistoryRepository struct {
	db                *gorm.DB
	singleflightGroup singleflight.Group
	redisCache        *cache.RedisCache
	logger            *logger.Logger
}

func NewOrderHistoryRepository(db *gorm.DB, redisClient *redis.Client, logger *logger.Logger) r.OrderHistoryRepository {
	redisCache := cache.NewRedisCache(redisClient, logger)
	return &orderHistoryRepository{
		db:         db,
		redisCache: redisCache,
		logger:     logger,
	}
}

type orderItemRepository struct {
	db                *gorm.DB
	singleflightGroup singleflight.Group
	redisCache        *cache.RedisCache
	logger            *logger.Logger
}

func NewOrderItemRepository(db *gorm.DB, redisClient *redis.Client, logger *logger.Logger) r.OrderItemRepository {
	redisCache := cache.NewRedisCache(redisClient, logger)
	return &orderItemRepository{
		db:         db,
		redisCache: redisCache,
		logger:     logger,
	}
}
