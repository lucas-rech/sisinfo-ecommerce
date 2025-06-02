package config

import (
	"fmt"

	"github.com/lucas-rech/sisinfo-ecommerce/internal/handler"
	"github.com/lucas-rech/sisinfo-ecommerce/internal/repository"
	"github.com/lucas-rech/sisinfo-ecommerce/internal/service"
	"gorm.io/gorm"
)

type Container struct {
	DB                *gorm.DB
	ProductRepository repository.ProductRepository
	UserRepository    repository.UserRepository

	ProductService service.ProductService
	UserService    service.UserService

	ProductHandler *handler.ProductHandler
}

func NewContainer() (*Container, error) {
	db, err := ConnectDatabase()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize database: %w", err)
	}

	userRepository := repository.NewUserRepository(db)
	productRepository := repository.NewProductRepository(db)

	userService := service.NewUserService(userRepository)
	productService := service.NewProductService(productRepository)

	productHandler := handler.NewProductHandler(productService)

	return &Container{
		DB:                db,
		ProductRepository: productRepository,
		UserRepository:    userRepository,
		ProductService:    productService,
		UserService:       userService,
		ProductHandler:    productHandler,
	}, nil

}
