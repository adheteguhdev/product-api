package user

import (
	"errors"
	"product-api/models"
	"time"
)

func (u *UserUsecase) GetAll() ([]*models.User, error) {
	return u.userRepo.GetAll()
}

func (u *UserUsecase) GetByID(id uint) (*models.User, error) {
	return u.userRepo.GetByID(id)
}

func (u *UserUsecase) Create(user *models.User) (*models.User, error) {
	if user.FullName == "" {
		msg := "full name is required"
		u.logger.Errorf(msg)
		return nil, errors.New(msg)
	}

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	return u.userRepo.Create(user)
}

func (u *UserUsecase) Update(user *models.User) (*models.User, error) {
	id := uint(user.ID)
	if id == 0 {
		msg := "user id is required"
		u.logger.Errorf(msg)
		return nil, errors.New(msg)
	}

	_, err := u.userRepo.GetByID(id)
	if err != nil {
		u.logger.Errorf(err.Error())
		return nil, err
	}

	if user.FullName == "" {
		msg := "full name is required"
		u.logger.Errorf(msg)
		return nil, errors.New(msg)
	}

	user.UpdatedAt = time.Now()
	return u.userRepo.Update(user)

}

func (u *UserUsecase) Delete(id uint) error {
	return u.userRepo.Delete(id)
}
