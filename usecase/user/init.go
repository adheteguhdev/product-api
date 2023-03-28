package user

import (
	"product-api/pkg/logger"
	"product-api/repository"
	uc "product-api/usecase"
)

type UserUsecase struct {
	userRepo repository.UserRepository
	logger   *logger.Logger
}

func NewUserUsecase(userRepo repository.UserRepository, logger *logger.Logger) uc.UserUsecase {
	return &UserUsecase{
		userRepo: userRepo,
		logger:   logger,
	}
}
