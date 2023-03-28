package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"product-api/models"
	"product-api/usecase"

	"github.com/labstack/echo/v4"
)

type orderItemHandler struct {
	orderItemUsecase usecase.OrderItemUsecase
}

func NewOrderItemHandler(e *echo.Echo, orderItemUsecase usecase.OrderItemUsecase) {
	handler := &orderItemHandler{
		orderItemUsecase: orderItemUsecase,
	}

	e.GET("/order", handler.GetAllOrderItem)
	e.GET("/order/:id", handler.GetOrderItemByID)
	e.POST("/order", handler.CreateOrderItem)
	e.PUT("/order/:id", handler.UpdateOrderItem)
	e.DELETE("/order/:id", handler.DeleteOrderItem)

}

func (u *orderItemHandler) GetAllOrderItem(c echo.Context) error {
	orderItems, err := u.orderItemUsecase.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "failed to get order items"})
	}
	return c.JSON(http.StatusOK, orderItems)
}

func (u *orderItemHandler) GetOrderItemByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid order item id"})
	}

	orderItem, err := u.orderItemUsecase.GetByID(uint(id))
	if err != nil {
		if err.Error() == "record not found" {
			return c.JSON(http.StatusNotFound, map[string]string{"message": "data not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "failed to get order item"})
	}
	return c.JSON(http.StatusOK, orderItem)
}

func (u *orderItemHandler) CreateOrderItem(c echo.Context) error {
	var orderItem models.OrderItem
	err := json.NewDecoder(c.Request().Body).Decode(&orderItem)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid request payload"})
	}

	createdOrderItem, err := u.orderItemUsecase.Create(&orderItem)
	if err != nil {
		if err.Error() == "name is required" {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
		}
		if err.Error() == "price is required" {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
		}
		if err.Error() == "duplicated key not allowed" {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": "name already exists"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "failed to create order item"})
	}
	return c.JSON(http.StatusCreated, createdOrderItem)
}

func (u *orderItemHandler) UpdateOrderItem(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid order item id"})
	}

	var orderItem models.OrderItem
	err = json.NewDecoder(c.Request().Body).Decode(&orderItem)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid request payload"})
	}

	orderItem.ID = uint64(id)

	_, err = u.orderItemUsecase.Update(&orderItem)
	if err != nil {
		if err.Error() == "record not found" {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": "data not found"})
		}
		if err.Error() == "name is required" {
			return c.JSON(http.StatusBadGateway, map[string]string{"message": err.Error()})
		}
		if err.Error() == "price is required" {
			return c.JSON(http.StatusBadGateway, map[string]string{"message": err.Error()})
		}
		if err.Error() == "duplicated key not allowed" {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": "name already exists"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "failed to update order item"})
	}
	return c.JSON(http.StatusOK, orderItem)
}

func (u *orderItemHandler) DeleteOrderItem(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid order item id"})
	}

	err = u.orderItemUsecase.Delete(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "failed to delete order item"})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "deleted successfully!"})
}
