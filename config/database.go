package config

import (
	"log"
	"os"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/lucas-rech/sisinfo-ecommerce/internal/domain"
)

var DB *gorm.DB

func ConnectDatabase() {
	LoadEnv()

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Sao_Paulo",
		host, user, password, dbname, port)
	
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	DB = db
	if err := DB.AutoMigrate(
		&domain.Cart{},
		&domain.CartItem{},
		&domain.Product{},
		&domain.User{},
		&domain.Order{},
		&domain.OrdemItem{},
	
	); err != nil {
		log.Fatal("Failed to migrate database schema:", err)
	}

	log.Println("Database connection established successfully")
}

