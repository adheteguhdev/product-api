package config

import (
	"fmt"
	"os"
	"product-api/models"
	"product-api/pkg/logger"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB

func InitDatabase() {
	log := logger.NewLogger()

	err := godotenv.Load()
	if err != nil {
		log.Errorf("failed to load .env file")
	}

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", host, user, password, dbName, port)
	database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	if err := database.AutoMigrate(&models.User{}, &models.OrderItem{}, &models.OrderHistory{}); err != nil {
		log.Errorf("failed to auto migrate models: %v", err)
	}

}

func DB() *gorm.DB {
	return database
}
