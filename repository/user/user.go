package user

import (
	"context"
	"fmt"
	"product-api/models"
	"strconv"
	"time"
)

func (r *userRepository) Create(user *models.User) (*models.User, error) {
	val, err, _ := r.singleflightGroup.Do("create", func() (interface{}, error) {
		result := r.db.Create(user)
		if result.Error != nil {
			msg := fmt.Sprintf("%v", result.Error)
			r.logger.Errorf(msg)
			return nil, result.Error
		}

		r.logger.Infof("user created")
		return user, nil
	})

	if err != nil {
		r.logger.Errorf(err.Error())
		return nil, err
	}

	keys := fmt.Sprintf("user:%d", user.ID)
	err = r.redisCache.SetEX(context.Background(), keys, user, time.Minute*5)
	if err != nil {
		r.logger.Errorf("failed to set user in Redis cache: %v", err)
	}

	return val.(*models.User), nil
}

func (r *userRepository) GetAll() ([]*models.User, error) {
	var users []*models.User

	cacheKey := "users"
	err := r.redisCache.Get(context.Background(), cacheKey, &users)
	if err == nil {
		r.logger.Infof("got all users from Redis cache")
		return users, nil
	}

	val, err, _ := r.singleflightGroup.Do("all", func() (interface{}, error) {
		result := r.db.Find(&users)
		if result.Error != nil {
			msg := fmt.Sprintf("%v", result.Error)
			r.logger.Errorf(msg)
			return nil, result.Error
		}
		r.logger.Infof("get all users")
		return users, nil
	})

	if err != nil {
		r.logger.Errorf(err.Error())
		return nil, err
	}

	err = r.redisCache.SetEX(context.Background(), cacheKey, users, time.Minute*5)
	if err != nil {
		r.logger.Errorf("failed to set all users in Redis cache: %v", err)
	}

	return val.([]*models.User), nil
}

func (r *userRepository) GetByID(id uint) (*models.User, error) {
	user := &models.User{}

	cacheKey := fmt.Sprintf("user:%d", id)
	err := r.redisCache.Get(context.Background(), cacheKey, user)

	if err == nil {
		r.logger.Infof("got user %d from Redis cache", id)
		return user, nil
	}

	val, err, _ := r.singleflightGroup.Do(strconv.Itoa(int(id)), func() (interface{}, error) {
		result := r.db.First(&user, id)

		if result.Error != nil {
			msg := fmt.Sprintf("%v", result.Error)
			r.logger.Errorf(msg)
			return nil, result.Error
		}

		r.logger.Infof("get user by id")
		return user, nil
	})

	if err != nil {
		r.logger.Errorf(err.Error())
		return nil, err
	}

	err = r.redisCache.SetEX(context.Background(), cacheKey, user, time.Minute*5)
	if err != nil {
		r.logger.Errorf("failed to set user %d in Redis cache: %v", id, err)
	}

	return val.(*models.User), nil
}

func (r *userRepository) Update(user *models.User) (*models.User, error) {
	result := r.db.Omit("created_at").Save(user)
	if result.Error != nil {
		msg := fmt.Sprintf("%v", result.Error)
		r.logger.Errorf((msg))
		return nil, result.Error
	}

	cacheKey := fmt.Sprintf("user:%d", user.ID)

	err := r.redisCache.Delete(context.Background(), cacheKey)
	if err != nil {
		r.logger.Errorf("error deleting user from Redis cache:", err)
	}

	r.logger.Infof("user updated")
	return user, nil
}

func (r *userRepository) Delete(id uint) error {
	user := &models.User{}

	cacheKey := fmt.Sprintf("user:%d", id)
	err := r.redisCache.Get(context.Background(), cacheKey, user)
	if err == nil {
		err := r.redisCache.Delete(context.Background(), cacheKey)
		if err != nil {
			r.logger.Errorf("error deleting user from Redis cache: %v", err)
		}
	}

	result := r.db.Delete(&models.User{}, id)
	if result.Error != nil {
		msg := fmt.Sprintf("%v", result.Error)
		r.logger.Errorf(msg)
		return result.Error
	}

	r.logger.Infof("user deleted")
	return nil
}
