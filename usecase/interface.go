package usecase

import "product-api/models"

type UserUsecase interface {
	Create(user *models.User) (*models.User, error)
	GetByID(id uint) (*models.User, error)
	Update(user *models.User) (*models.User, error)
	Delete(id uint) error
	GetAll() ([]*models.User, error)
}

type OrderItemUsecase interface {
	Create(orderItem *models.OrderItem) (*models.OrderItem, error)
	GetByID(id uint) (*models.OrderItem, error)
	Update(orderItem *models.OrderItem) (*models.OrderItem, error)
	Delete(id uint) error
	GetAll() ([]*models.OrderItem, error)
}

type OrderHistoryUsecase interface {
	Create(orderHistory *models.OrderHistory) (*models.OrderHistory, error)
	GetByID(id uint) (*models.OrderHistory, error)
	Update(orderHistory *models.OrderHistory) (*models.OrderHistory, error)
	GetAll() ([]*models.OrderHistory, error)
}
