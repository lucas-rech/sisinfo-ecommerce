package repository

import "github.com/lucas-rech/sisinfo-ecommerce/internal/domain"

type UserRepository interface {
	Create(user *domain.User) error
	FindByID(id uint) (*domain.User, error)
	FindByEmail(email string) (*domain.User, error)
	Update(user *domain.User) error
	Delete(id uint) error
}
