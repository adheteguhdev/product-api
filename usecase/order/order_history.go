package order

import (
	"errors"
	"product-api/models"
	"time"
)

func (u *OrderHistoryUsecase) GetAll() ([]*models.OrderHistory, error) {
	return u.orderHistoryRepo.GetAll()
}

func (u *OrderHistoryUsecase) GetByID(id uint) (*models.OrderHistory, error) {
	return u.orderHistoryRepo.GetByID(id)
}

func (u *OrderHistoryUsecase) Create(orderHistory *models.OrderHistory) (*models.OrderHistory, error) {
	if orderHistory.UserID == 0 {
		msg := "user id is required"
		u.logger.Errorf(msg)
		return nil, errors.New(msg)
	}
	if orderHistory.OrderItemID == 0 {
		msg := "order item id is required"
		u.logger.Errorf(msg)
		return nil, errors.New(msg)
	}

	userID := orderHistory.UserID
	userExists, err := u.userRepo.GetByID(uint(userID))
	if err != nil {
		msg := "user id not found"
		u.logger.Errorf(msg)
		return nil, errors.New(msg)
	}
	if userExists == nil {
		msg := "user not found"
		u.logger.Errorf(msg)
		return nil, errors.New(msg)
	}

	orderItemId := orderHistory.OrderItemID
	orderItemExists, err := u.orderItemRepo.GetByID(uint(orderItemId))
	if err != nil {
		msg := "order item id not found"
		u.logger.Errorf(msg)
		return nil, errors.New(msg)
	}
	if orderItemExists == nil {
		msg := "order item not found"
		u.logger.Errorf(msg)
		return nil, errors.New(msg)
	}
	orderHistory.CreatedAt = time.Now()
	orderHistory.UpdatedAt = time.Now()
	return u.orderHistoryRepo.Create(orderHistory)
}

func (u *OrderHistoryUsecase) Update(orderHistory *models.OrderHistory) (*models.OrderHistory, error) {
	id := uint(orderHistory.ID)
	if id == 0 {
		msg := "order history id is required"
		u.logger.Errorf(msg)
		return nil, errors.New(msg)
	}
	if orderHistory.UserID == 0 {
		msg := "user id is required"
		u.logger.Errorf(msg)
		return nil, errors.New(msg)
	}
	if orderHistory.OrderItemID == 0 {
		msg := "order id is required"
		u.logger.Errorf(msg)
		return nil, errors.New(msg)
	}
	_, err := u.orderHistoryRepo.GetByID(id)
	if err != nil {
		u.logger.Errorf(err.Error())
		return nil, err
	}

	userID := orderHistory.UserID
	userExists, err := u.userRepo.GetByID(uint(userID))
	if err != nil {
		msg := "user id not found"
		u.logger.Errorf(msg)
		return nil, errors.New(msg)
	}
	if userExists == nil {
		msg := "user not found"
		u.logger.Errorf(msg)
		return nil, errors.New(msg)
	}

	orderItemId := orderHistory.OrderItemID
	orderItemExists, err := u.orderItemRepo.GetByID(uint(orderItemId))
	if err != nil {
		msg := "order item id not found"
		u.logger.Errorf(msg)
		return nil, errors.New(msg)
	}
	if orderItemExists == nil {
		msg := "order item not found"
		u.logger.Errorf(msg)
		return nil, errors.New(msg)
	}

	orderHistory.UpdatedAt = time.Now()
	return u.orderHistoryRepo.Update(orderHistory)
}
