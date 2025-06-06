package service

import (
	"github.com/lucas-rech/sisinfo-ecommerce/backend/internal/dto"
)

type UserService interface {
	CreateUser(user dto.UserCreateRequest) error
	FindUserByID(id uint) (*dto.UserResponse, error)
	FindUserIDByEmail(email string) (uint, error)
	FindUserByEmail(email string) (*dto.UserResponse, error)
	UpdateUser(user dto.UserUpdateRequest, id uint) error
	DeleteUser(id uint) error
	Login(email, password string) (*dto.UserResponse, error)
}

type ProductService interface {
	CreateProduct(product dto.ProductCreateRequest) error
	FindProductByID(id uint) (*dto.ProductResponse, error)
	FindAllProducts() ([]dto.ProductResponse, error)
	UpdateProduct(product dto.ProductUpdateRequest, id *uint) error
	DeleteProduct(id uint) error
	IncreaseProductStock(id uint, quantity int) error
	DecreaseProductStock(id uint, quantity int) error
}

type CartItemService interface {
	AddItemToCart(userID uint, productID uint, quantity int) error
	GetItemsByCartID(cartID uint) ([]dto.CartItemResponse, error)
	RemoveItemFromCart(item *dto.CartItemDeleteRequest) error
	UpdateItemInCart(request dto.CartItemUpdateRequest, userID uint) error
}