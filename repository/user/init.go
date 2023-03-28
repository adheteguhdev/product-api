package user

import (
	"product-api/pkg/redis/cache"
	r "product-api/repository"

	"product-api/pkg/logger"

	"github.com/go-redis/redis/v8"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
)

type userRepository struct {
	db                *gorm.DB
	singleflightGroup singleflight.Group
	redisCache        *cache.RedisCache
	logger            *logger.Logger
}

func NewUserRepository(db *gorm.DB, redisClient *redis.Client, logger *logger.Logger) r.UserRepository {
	redisCache := cache.NewRedisCache(redisClient, logger)
	return &userRepository{
		db:         db,
		redisCache: redisCache,
		logger:     logger,
	}
}
