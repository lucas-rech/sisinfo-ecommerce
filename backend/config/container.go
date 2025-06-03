package config

import (
	"fmt"

	"github.com/lucas-rech/sisinfo-ecommerce/backend/internal/handler"
	"github.com/lucas-rech/sisinfo-ecommerce/backend/internal/repository"
	"github.com/lucas-rech/sisinfo-ecommerce/backend/internal/service"
	"gorm.io/gorm"
)

type Container struct {
	DB                *gorm.DB
	ProductRepository repository.ProductRepository
	UserRepository    repository.UserRepository
	CartRepository    repository.CartRepository
	CartItemRepository repository.CartItemRepository

	ProductService service.ProductService
	UserService    service.UserService

	ProductHandler *handler.ProductHandler
	UserHandler    *handler.UserHandler
}

func NewContainer() (*Container, error) {
	db, err := ConnectDatabase()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize database: %w", err)
	}

	userRepository := repository.NewUserRepository(db)
	cartRepository := repository.NewCartRepository(db)
	cartItemRepository := repository.NewCartItemRepository(db)
	productRepository := repository.NewProductRepository(db)

	userService := service.NewUserService(userRepository, cartRepository)
	productService := service.NewProductService(productRepository)

	productHandler := handler.NewProductHandler(productService)
	userHandler := handler.NewUserHandler(userService)

	return &Container{
		DB:                db,
		ProductRepository: productRepository,
		UserRepository:    userRepository,
		ProductService:    productService,
		UserService:       userService,
		ProductHandler:    productHandler,
		UserHandler:       userHandler,
		CartRepository:    cartRepository,
		CartItemRepository: cartItemRepository,
	}, nil

}
