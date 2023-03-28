package user

import (
	"product-api/models"
	"product-api/pkg/logger"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAll(t *testing.T) {

	userRepo := new(mockUserRepository)
	logger := logger.NewLogger()

	userUsecase := NewUserUsecase(userRepo, logger)

	users := []*models.User{
		{ID: 1, FullName: "John Doe"},
		{ID: 2, FullName: "Jane Doe"},
	}

	userRepo.On("GetAll").Return(users, nil)

	result, err := userUsecase.GetAll()

	assert.NoError(t, err)
	assert.Equal(t, users, result)

	userRepo.AssertExpectations(t)
}

type mockUserRepository struct {
	mock.Mock
}

func (m *mockUserRepository) GetAll() ([]*models.User, error) {
	args := m.Called()
	return args.Get(0).([]*models.User), args.Error(1)
}

func (m *mockUserRepository) GetByID(id uint) (*models.User, error) {
	args := m.Called(id)
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *mockUserRepository) Create(user *models.User) (*models.User, error) {
	args := m.Called(user)
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *mockUserRepository) Update(user *models.User) (*models.User, error) {
	args := m.Called(user)
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *mockUserRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}
