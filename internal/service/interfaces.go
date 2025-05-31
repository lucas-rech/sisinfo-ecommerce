package service

import (
	"github.com/lucas-rech/sisinfo-ecommerce/internal/domain"
	"github.com/lucas-rech/sisinfo-ecommerce/internal/dto"
)

type UserService interface {
	CreateUser(user dto.UserCreateRequest) error
	FindUserByID(id uint) (*domain.User, error)
	FindUserByEmail(email string) (*domain.User, error)
	UpdateUser(user dto.UserUpdateRequest, id uint) error
	DeleteUser(id uint) error
}

type ProductService interface {
	CreateProduct(product dto.ProductCreateRequest) error
	FindProductByID(id uint) (*domain.Product, error)
	FindAllProducts() ([]domain.Product, error)
	UpdateProduct(product dto.ProductUpdateRequest, id *uint) error
	DeleteProduct(id uint) error
}
