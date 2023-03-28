package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"product-api/models"
	"product-api/usecase"

	"github.com/labstack/echo/v4"
)

type userHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(e *echo.Echo, userUsecase usecase.UserUsecase) {
	handler := &userHandler{
		userUsecase: userUsecase,
	}

	e.GET("/user", handler.GetAllUsers)
	e.GET("/user/:id", handler.GetUserByID)
	e.POST("/user", handler.CreateUser)
	e.PUT("/user/:id", handler.UpdateUser)
	e.DELETE("/user/:id", handler.DeleteUser)

}

func (u *userHandler) GetAllUsers(c echo.Context) error {
	users, err := u.userUsecase.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "failed to get users"})
	}
	return c.JSON(http.StatusOK, users)
}

func (u *userHandler) GetUserByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid user id"})
	}

	user, err := u.userUsecase.GetByID(uint(id))
	if err != nil {
		if err.Error() == "record not found" {
			return c.JSON(http.StatusNotFound, map[string]string{"message": "data not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "failed to get user"})
	}
	return c.JSON(http.StatusOK, user)
}

func (u *userHandler) CreateUser(c echo.Context) error {
	var user models.User
	err := json.NewDecoder(c.Request().Body).Decode(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid request payload"})
	}

	createdUser, err := u.userUsecase.Create(&user)
	if err != nil {
		if err.Error() == "full name is required" {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "failed to create user"})
	}
	return c.JSON(http.StatusCreated, createdUser)
}

func (u *userHandler) UpdateUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid user id"})
	}

	var user models.User
	err = json.NewDecoder(c.Request().Body).Decode(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid request payload"})
	}

	user.ID = uint64(id)

	_, err = u.userUsecase.Update(&user)
	if err != nil {
		if err.Error() == "record not found" {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": "data not found"})
		}
		if err.Error() == "full name is required" {
			return c.JSON(http.StatusBadGateway, map[string]string{"message": err.Error()})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "failed to update user"})
	}
	return c.JSON(http.StatusOK, user)
}

func (u *userHandler) DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid user id"})
	}

	err = u.userUsecase.Delete(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "failed to delete user"})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "deleted successfully!"})
}
