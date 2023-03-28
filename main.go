package main

import (
	"fmt"
	"os"
	"product-api/config"
	"product-api/handler"
	"product-api/pkg/logger"
	rOrder "product-api/repository/order"
	rUser "product-api/repository/user"
	uOrder "product-api/usecase/order"
	uUser "product-api/usecase/user"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	log := logger.NewLogger()

	err := godotenv.Load()
	if err != nil {
		log.Errorf("failed to load .env file error: %v", err)
		panic(err)
	}

	e := echo.New()

	config.InitDatabase()
	db := config.DB()

	dbGorm, err := db.DB()
	if err != nil {
		log.Errorf("error initializing database connection: %v", err)
		panic(err)
	}

	dbGorm.Ping()

	log.Infof("Database connected")

	err = config.InitRedis()
	if err != nil {
		log.Errorf("error initializing redis connection: %v", err)
	}

	userRepo := rUser.NewUserRepository(db, config.RedisClient, log)
	userUsecase := uUser.NewUserUsecase(userRepo, log)
	orderItemRepo := rOrder.NewOrderItemRepository(db, config.RedisClient, log)
	orderItemUsecase := uOrder.NewOrderItemUsecase(orderItemRepo, log)
	orderHistoryRepo := rOrder.NewOrderHistoryRepository(db, config.RedisClient, log)
	orderHistoryUsecase := uOrder.NewOrderHistoryUsecase(orderHistoryRepo, userRepo, orderItemRepo, log)

	handler.NewUserHandler(e, userUsecase)
	handler.NewOrderItemHandler(e, orderItemUsecase)
	handler.NewOrderHistoryHandler(e, orderHistoryUsecase)

	port := os.Getenv("APP_PORT")
	log.Infof("Starting server on port %s", port)

	addr := fmt.Sprintf(":%s", port)
	if err := e.Start(addr); err != nil {
		log.Errorf("Error starting server: %v", err)
	}
}
