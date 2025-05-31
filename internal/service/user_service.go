package service

import (
	"fmt"

	"github.com/lucas-rech/sisinfo-ecommerce/internal/domain"
	"github.com/lucas-rech/sisinfo-ecommerce/internal/repository"
	"gorm.io/gorm"
)

//Análogo a uma classe
type UserService struct {
	userRepo repository.UserRepository
}

//Análogo ao construtor
// NewUserService cria uma nova instância de UserService
func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) CreateUser(user *domain.User) error {
	if user.Email == "" || user.Password == "" {
		return fmt.Errorf("email and password are required")
	}

	existingUser, err := s.userRepo.FindByEmail(user.Email)
	if err != nil && err != gorm.ErrRecordNotFound {
		return fmt.Errorf("error checking for existing user: %w", err)
	}
	if existingUser != nil {
		return fmt.Errorf("user with this email already exists")
	}
	
	return s.userRepo.Create(user)
}








