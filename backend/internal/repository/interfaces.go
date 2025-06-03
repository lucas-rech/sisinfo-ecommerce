package repository

import "github.com/lucas-rech/sisinfo-ecommerce/backend/internal/domain"

type UserRepository interface {
	Create(user *domain.User) error
	FindByID(id *uint) (*domain.User, error)
	FindByEmail(email string) (*domain.User, error)
	Update(user *domain.User) error
	Delete(id *uint) error
}

type ProductRepository interface {
	Create(product *domain.Product) error
	FindByID(id *uint) (*domain.Product, error)
	FindAll() ([]domain.Product, error)
	Update(product *domain.Product) error
	Delete(id *uint) error
}
