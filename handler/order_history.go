package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"product-api/models"
	"product-api/usecase"

	"github.com/labstack/echo/v4"
)

type orderHistoryHandler struct {
	orderHistoryUsecase usecase.OrderHistoryUsecase
}

func NewOrderHistoryHandler(e *echo.Echo, orderHistoryUsecase usecase.OrderHistoryUsecase) {
	handler := &orderHistoryHandler{
		orderHistoryUsecase: orderHistoryUsecase,
	}

	e.GET("/order-history", handler.GetAllOrderHistory)
	e.GET("/order-history/:id", handler.GetOrderHistory)
	e.POST("/order-history", handler.CreateOrderHistory)
	e.PUT("/order-history/:id", handler.UpdateOrderHistory)

}

func (u *orderHistoryHandler) GetAllOrderHistory(c echo.Context) error {
	orderHistories, err := u.orderHistoryUsecase.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "failed to get order items"})
	}
	return c.JSON(http.StatusOK, orderHistories)
}

func (u *orderHistoryHandler) GetOrderHistory(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid order history id"})
	}

	orderHistory, err := u.orderHistoryUsecase.GetByID(uint(id))
	if err != nil {
		if err.Error() == "record not found" {
			return c.JSON(http.StatusNotFound, map[string]string{"message": "data not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "failed to get order item"})
	}
	return c.JSON(http.StatusOK, orderHistory)
}

func (u *orderHistoryHandler) CreateOrderHistory(c echo.Context) error {
	var orderHistory models.OrderHistory
	err := json.NewDecoder(c.Request().Body).Decode(&orderHistory)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid request payload"})
	}

	createdOrderHistory, err := u.orderHistoryUsecase.Create(&orderHistory)
	if err != nil {
		if err.Error() == "user id is required" {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
		}
		if err.Error() == "order item id is required" {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
		}
		if err.Error() == "user id not found" {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
		}
		if err.Error() == "order item id not found" {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
		}

		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "failed to create order item"})
	}
	return c.JSON(http.StatusCreated, createdOrderHistory)
}

func (u *orderHistoryHandler) UpdateOrderHistory(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid order history id"})
	}

	var orderHistory models.OrderHistory
	err = json.NewDecoder(c.Request().Body).Decode(&orderHistory)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid request payload"})
	}

	orderHistory.ID = uint64(id)

	_, err = u.orderHistoryUsecase.Update(&orderHistory)
	if err != nil {
		if err.Error() == "record not found" {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": "data not found"})
		}
		if err.Error() == "user id is required" {
			return c.JSON(http.StatusBadGateway, map[string]string{"message": err.Error()})
		}
		if err.Error() == "order item id is required" {
			return c.JSON(http.StatusBadGateway, map[string]string{"message": err.Error()})
		}
		if err.Error() == "user id not found" {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
		}
		if err.Error() == "order item id not found" {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
		}

		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "failed to update order item"})
	}
	return c.JSON(http.StatusOK, orderHistory)
}
