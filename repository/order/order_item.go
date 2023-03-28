package order

import (
	"context"
	"fmt"
	"product-api/models"
	"strconv"
	"time"
)

func (r *orderItemRepository) Create(orderItem *models.OrderItem) (*models.OrderItem, error) {
	val, err, _ := r.singleflightGroup.Do("create", func() (interface{}, error) {
		result := r.db.Create(orderItem)
		if result.Error != nil {
			msg := fmt.Sprintf("%v", result.Error)
			r.logger.Errorf(msg)
			return nil, result.Error
		}
		r.logger.Infof("order created")
		return orderItem, nil
	})

	if err != nil {
		r.logger.Errorf(err.Error())
		return nil, err
	}

	err = r.redisCache.SetEX(context.Background(), fmt.Sprintf("orderItem:%d", orderItem.ID), orderItem, time.Minute*5)
	if err != nil {
		r.logger.Errorf("failed to set order item in Redis cache: %v", err)
	}

	return val.(*models.OrderItem), nil
}

func (r *orderItemRepository) GetAll() ([]*models.OrderItem, error) {
	var orderItem []*models.OrderItem

	cacheKey := "orderItems"
	err := r.redisCache.Get(context.Background(), cacheKey, &orderItem)
	if err == nil {
		r.logger.Infof("got all order items from Redis cache")
		return orderItem, nil
	}

	val, err, _ := r.singleflightGroup.Do("all", func() (interface{}, error) {
		result := r.db.Find(&orderItem)
		if result.Error != nil {
			msg := fmt.Sprintf("%v", result.Error)
			r.logger.Errorf(msg)
			return nil, result.Error
		}

		r.logger.Infof("get all order items")
		return orderItem, nil
	})

	if err != nil {
		r.logger.Errorf(err.Error())
		return nil, err
	}

	err = r.redisCache.SetEX(context.Background(), cacheKey, orderItem, time.Minute*5)
	if err != nil {
		r.logger.Errorf("failed to set all order items in Redis cache: %v", err)
	}

	return val.([]*models.OrderItem), nil
}

func (r *orderItemRepository) GetByID(id uint) (*models.OrderItem, error) {
	orderItem := &models.OrderItem{}

	cacheKey := fmt.Sprintf("orderItem:%d", id)
	err := r.redisCache.Get(context.Background(), cacheKey, orderItem)

	if err == nil {
		r.logger.Infof("got orderItem %d from Redis cache", id)
		return orderItem, nil
	}

	val, err, _ := r.singleflightGroup.Do(strconv.Itoa(int(id)), func() (interface{}, error) {
		result := r.db.First(&orderItem, id)

		if result.Error != nil {
			msg := fmt.Sprintf("%v", result.Error)
			r.logger.Errorf(msg)
			return nil, result.Error
		}

		r.logger.Infof("get order item by id")
		return orderItem, nil
	})

	if err != nil {
		r.logger.Errorf(err.Error())
		return nil, err
	}

	err = r.redisCache.SetEX(context.Background(), cacheKey, orderItem, time.Minute*5)
	if err != nil {
		r.logger.Errorf("failed to set orderItem %d in Redis cache: %v", id, err)
	}

	return val.(*models.OrderItem), nil
}

func (r *orderItemRepository) Update(orderItem *models.OrderItem) (*models.OrderItem, error) {
	result := r.db.Omit("created_at").Save(orderItem)
	if result.Error != nil {
		msg := fmt.Sprintf("%v", result.Error)
		r.logger.Errorf((msg))
		return nil, result.Error
	}

	cacheKey := fmt.Sprintf("orderItem:%d", orderItem.ID)

	err := r.redisCache.Delete(context.Background(), cacheKey)
	if err != nil {
		r.logger.Errorf("error deleting orderItem from Redis cache:", err)
	}

	r.logger.Infof("order item updated")
	return orderItem, nil
}

func (r *orderItemRepository) Delete(id uint) error {
	orderItem := &models.OrderItem{}

	cacheKey := fmt.Sprintf("orderItem:%d", id)
	err := r.redisCache.Get(context.Background(), cacheKey, orderItem)
	if err == nil {
		err := r.redisCache.Delete(context.Background(), cacheKey)
		if err != nil {
			r.logger.Errorf("error deleting orderItem from Redis cache: %v", err)
		}
	}

	result := r.db.Delete(&models.OrderItem{}, id)
	if result.Error != nil {
		return result.Error
	}

	r.logger.Infof("order item deleted")
	return nil
}
