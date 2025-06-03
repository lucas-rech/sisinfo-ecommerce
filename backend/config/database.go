package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/lucas-rech/sisinfo-ecommerce/backend/internal/domain"
	"github.com/lucas-rech/sisinfo-ecommerce/backend/utils"
)

func ConnectDatabase() (*gorm.DB, error) {
	utils.LoadEnv()

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Sao_Paulo",
		host, user, password, dbname, port)
	
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	if err := db.AutoMigrate(
		&domain.Cart{},
		&domain.CartItem{},
		&domain.Product{},
		&domain.User{},
		&domain.Order{},
		&domain.OrderItem{},
	
	); err != nil {
		return nil, fmt.Errorf("failed to auto migrate database: %w", err)
	}

	log.Println("Database connection established successfully")
	return db, nil
}

