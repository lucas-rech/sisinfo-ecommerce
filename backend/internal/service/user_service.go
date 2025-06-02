package service

import (
	"fmt"
	"time"

	"github.com/lucas-rech/sisinfo-ecommerce/internal/domain"
	"github.com/lucas-rech/sisinfo-ecommerce/internal/dto"
	"github.com/lucas-rech/sisinfo-ecommerce/internal/repository"
	"github.com/lucas-rech/sisinfo-ecommerce/pkg/utils"
	"gorm.io/gorm"
)

// Análogo a uma classe que implementa a interface UserService
type userService struct {
	userRepo repository.UserRepository
}

// Análogo ao construtor
// NewUserService cria uma nova instância de UserService
func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) CreateUser(user dto.UserCreateRequest) error {
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

	hashedPassword, err := utils.HashPassword(&user.Password)
	if err != nil {
		return err
	}

	userEntity := domain.User{
		Name:      user.Name,
		Email:     user.Email,
		Password:  hashedPassword,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Role:      domain.RoleCustomer, // Por padrão, define o papel como cliente
	}

	return s.userRepo.Create(&userEntity)
}

func (s *userService) FindUserByID(id uint) (*dto.UserResponse, error) {
	if id == 0 {
		return nil, fmt.Errorf("invalid user ID")
	}
	user, err := s.userRepo.FindByID(&id)
	if err != nil {
		return nil, fmt.Errorf("error finding user by ID: %w", err)
	}
	if user == nil {
		return nil, fmt.Errorf("user not found")
	}

	userResponse := dto.UserResponse{
		Name:  user.Name,
		Email: user.Email,
		Role:  string(user.Role),
	}
	return &userResponse, nil
}

func (s *userService) FindUserByEmail(email string) (*dto.UserResponse, error) {
	if email == "" {
		return nil, fmt.Errorf("email is required")
	}

	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("error finding user by email: %w", err)
	}
	if user == nil {
		return nil, fmt.Errorf("user not found")
	}

	userResponse := dto.UserResponse{
		Name:  user.Name,
		Email: user.Email,
		Role:  string(user.Role),
	}

	return &userResponse, nil
}

// FETCH Update
func (s *userService) UpdateUser(request dto.UserUpdateRequest, id uint) error {
	if id == 0 {
		return fmt.Errorf("invalid user ID")
	}

	existingUser, err := s.userRepo.FindByID(&id)
	if err != nil {
		return fmt.Errorf("error finding user by ID: %w", err)
	}
	if existingUser == nil {
		return fmt.Errorf("user not found")
	}

	// Atualiza somente os campos que foram fornecidos no request
	if request.Name != nil {
		existingUser.Name = *request.Name
	}
	if request.Email != nil {
		existingUser.Email = *request.Email
	}

	// Se a senha for fornecida, atualiza o campo de senha
	if request.Password != nil {
		hashedPassword, err := utils.HashPassword(request.Password)
		if err != nil {
			return err
		}
		existingUser.Password = hashedPassword
	}

	return s.userRepo.Update(existingUser)
}

func (s *userService) DeleteUser(id uint) error {
	if id == 0 {
		return fmt.Errorf("invalid user ID")
	}

	existingUser, err := s.userRepo.FindByID(&id)
	if err != nil {
		return fmt.Errorf("error finding user by ID: %w", err)
	}
	if existingUser == nil {
		return fmt.Errorf("user not found")
	}

	return s.userRepo.Delete(&id)
}


