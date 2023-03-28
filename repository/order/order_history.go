package order

import (
	"context"
	"fmt"
	"product-api/models"
	"strconv"
	"time"
)

func (r *orderHistoryRepository) Create(orderHistory *models.OrderHistory) (*models.OrderHistory, error) {
	val, err, _ := r.singleflightGroup.Do("create", func() (interface{}, error) {

		result := r.db.Create(orderHistory)
		r.db.Preload("User").Preload("OrderItem").Find(orderHistory)
		if result.Error != nil {
			msg := fmt.Sprintf("%v", result.Error)
			r.logger.Errorf(msg)
			return nil, result.Error
		}

		r.logger.Infof("order history created")
		return orderHistory, nil
	})

	if err != nil {
		r.logger.Errorf(err.Error())
		return nil, err
	}

	keys := fmt.Sprintf("orderHistory:%d", orderHistory.ID)
	err = r.redisCache.SetEX(context.Background(), keys, orderHistory, time.Minute*5)
	if err != nil {
		r.logger.Errorf("failed to set orderHistory in Redis cache: %v", err)
	}

	return val.(*models.OrderHistory), nil
}

func (r *orderHistoryRepository) GetAll() ([]*models.OrderHistory, error) {
	var orderHistory []*models.OrderHistory

	cacheKey := "orderHistory"
	err := r.redisCache.Get(context.Background(), cacheKey, &orderHistory)
	if err == nil {
		r.logger.Infof("got all orderHistory from Redis cache")
		return orderHistory, nil
	}

	val, err, _ := r.singleflightGroup.Do("all", func() (interface{}, error) {
		result := r.db.Preload("OrderItem").Preload("User").Find(&orderHistory)
		if result.Error != nil {
			msg := fmt.Sprintf("%v", result.Error)
			r.logger.Errorf(msg)
			return nil, result.Error
		}

		r.logger.Infof("get all order history")
		return orderHistory, nil
	})

	if err != nil {
		r.logger.Errorf(err.Error())
		return nil, err
	}

	err = r.redisCache.SetEX(context.Background(), cacheKey, orderHistory, time.Minute*5)
	if err != nil {
		r.logger.Errorf("failed to set all orderHistory in Redis cache: %v", err)
	}

	return val.([]*models.OrderHistory), nil
}

func (r *orderHistoryRepository) GetByID(id uint) (*models.OrderHistory, error) {
	orderHistory := &models.OrderHistory{}

	cacheKey := fmt.Sprintf("orderHistory:%d", id)
	err := r.redisCache.Get(context.Background(), cacheKey, orderHistory)

	if err == nil {
		r.logger.Infof("got orderHistory %d from Redis cache", id)
		return orderHistory, nil
	}

	val, err, _ := r.singleflightGroup.Do(strconv.Itoa(int(id)), func() (interface{}, error) {
		result := r.db.Preload("OrderItem").Preload("User").First(&orderHistory, id)

		if result.Error != nil {
			msg := fmt.Sprintf("%v", result.Error)
			r.logger.Errorf(msg)
			return nil, result.Error
		}

		r.logger.Infof("get order history by id")
		return orderHistory, nil
	})

	if err != nil {
		r.logger.Errorf(err.Error())
		return nil, err
	}

	err = r.redisCache.SetEX(context.Background(), cacheKey, orderHistory, time.Minute*5)
	if err != nil {
		r.logger.Errorf("failed to set orderHistory %d in Redis cache: %v", id, err)
	}

	return val.(*models.OrderHistory), nil
}

func (r *orderHistoryRepository) Update(orderHistory *models.OrderHistory) (*models.OrderHistory, error) {
	result := r.db.Omit("created_at").Save(orderHistory)
	r.db.Preload("User").Preload("OrderItem").Find(orderHistory)
	if result.Error != nil {
		msg := fmt.Sprintf("%v", result.Error)
		r.logger.Errorf((msg))
		return nil, result.Error
	}

	cacheKey := fmt.Sprintf("orderHistory:%d", orderHistory.ID)

	err := r.redisCache.Delete(context.Background(), cacheKey)
	if err != nil {
		r.logger.Errorf("error deleting orderHistory from Redis cache:", err)
	}

	r.logger.Infof("order history updated")
	return orderHistory, nil
}
