package service

import (
	"github.com/lucas-rech/sisinfo-ecommerce/backend/internal/dto"
)

type UserService interface {
	CreateUser(user dto.UserCreateRequest) error
	FindUserByID(id uint) (*dto.UserResponse, error)
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
}
