package order

import (
	"product-api/pkg/logger"
	"product-api/repository"
	uc "product-api/usecase"
)

type OrderItemUsecase struct {
	orderItemRepo repository.OrderItemRepository
	logger        *logger.Logger
}

func NewOrderItemUsecase(orderItemRepo repository.OrderItemRepository, logger *logger.Logger) uc.OrderItemUsecase {
	return &OrderItemUsecase{
		orderItemRepo: orderItemRepo,
		logger:        logger,
	}
}

type OrderHistoryUsecase struct {
	orderHistoryRepo repository.OrderHistoryRepository
	userRepo         repository.UserRepository
	orderItemRepo    repository.OrderItemRepository
	logger           *logger.Logger
}

func NewOrderHistoryUsecase(
	orderHistoryRepo repository.OrderHistoryRepository,
	userRepo repository.UserRepository,
	orderItemRepo repository.OrderItemRepository,
	logger *logger.Logger) uc.OrderHistoryUsecase {
	return &OrderHistoryUsecase{
		orderHistoryRepo: orderHistoryRepo,
		userRepo:         userRepo,
		orderItemRepo:    orderItemRepo,
		logger:           logger,
	}
}
