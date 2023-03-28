package order

import (
	"errors"
	"product-api/models"
	"time"
)

func (u *OrderItemUsecase) GetAll() ([]*models.OrderItem, error) {
	return u.orderItemRepo.GetAll()
}

func (u *OrderItemUsecase) GetByID(id uint) (*models.OrderItem, error) {
	return u.orderItemRepo.GetByID(id)
}

func (u *OrderItemUsecase) Create(orderItem *models.OrderItem) (*models.OrderItem, error) {
	if orderItem.Name == "" {
		msg := "name is required"
		u.logger.Errorf(msg)
		return nil, errors.New(msg)
	}
	if orderItem.Price == 0.00 {
		msg := "price is required"
		u.logger.Errorf(msg)
		return nil, errors.New(msg)
	}

	orderItem.CreatedAt = time.Now()
	orderItem.UpdatedAt = time.Now()
	return u.orderItemRepo.Create(orderItem)
}

func (u *OrderItemUsecase) Update(orderItem *models.OrderItem) (*models.OrderItem, error) {
	id := uint(orderItem.ID)
	if id == 0 {
		msg := "orderItem id is required"
		u.logger.Errorf(msg)
		return nil, errors.New(msg)
	}

	_, err := u.orderItemRepo.GetByID(id)
	if err != nil {
		u.logger.Errorf(err.Error())
		return nil, err
	}

	if orderItem.Name == "" {
		msg := "name is required"
		u.logger.Errorf(msg)
		return nil, errors.New(msg)
	}
	if orderItem.Price == 0.00 {
		msg := "price is required"
		u.logger.Errorf(msg)
		return nil, errors.New(msg)
	}

	orderItem.UpdatedAt = time.Now()
	return u.orderItemRepo.Update(orderItem)

}

func (u *OrderItemUsecase) Delete(id uint) error {
	return u.orderItemRepo.Delete(id)
}
