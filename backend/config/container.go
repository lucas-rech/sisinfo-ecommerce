package config

import (
	"fmt"

	"github.com/lucas-rech/sisinfo-ecommerce/backend/internal/handler"
	"github.com/lucas-rech/sisinfo-ecommerce/backend/internal/repository"
	"github.com/lucas-rech/sisinfo-ecommerce/backend/internal/service"
	"gorm.io/gorm"
)

type Container struct {
	DB                 *gorm.DB
	ProductRepository  repository.ProductRepository
	UserRepository     repository.UserRepository
	CartRepository     repository.CartRepository
	CartItemRepository repository.CartItemRepository

	ProductService  service.ProductService
	UserService     service.UserService
	CartItemService service.CartItemService

	ProductHandler  *handler.ProductHandler
	UserHandler     *handler.UserHandler
	CartItemHandler *handler.CartItemHandler
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
	cartItemService := service.NewCartItemService(cartItemRepository, cartRepository, productService)

	productHandler := handler.NewProductHandler(productService)
	userHandler := handler.NewUserHandler(userService)
	cartItemHandler := handler.NewCartItemHandler(cartItemService, userService)

	return &Container{
		DB:                 db,
		ProductRepository:  productRepository,
		UserRepository:     userRepository,
		ProductService:     productService,
		CartItemService:    cartItemService,
		UserService:        userService,
		ProductHandler:     productHandler,
		UserHandler:        userHandler,
		CartRepository:     cartRepository,
		CartItemRepository: cartItemRepository,
		CartItemHandler:    cartItemHandler,
	}, nil

}
