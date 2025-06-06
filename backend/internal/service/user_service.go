package service

import (
	"fmt"
	"time"

	"github.com/lucas-rech/sisinfo-ecommerce/backend/internal/domain"
	"github.com/lucas-rech/sisinfo-ecommerce/backend/internal/dto"
	"github.com/lucas-rech/sisinfo-ecommerce/backend/internal/repository"
	"github.com/lucas-rech/sisinfo-ecommerce/backend/utils"
	"gorm.io/gorm"
)

// Análogo a uma classe que implementa a interface UserService
type userService struct {
	userRepo repository.UserRepository
	cartRepo repository.CartRepository
}

// Análogo ao construtor
// NewUserService cria uma nova instância de UserService
func NewUserService(userRepo repository.UserRepository, cartRepo repository.CartRepository) UserService {
	return &userService{
		userRepo: userRepo,
		cartRepo: cartRepo,
	}
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
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  hashedPassword,
		Phone:     user.Phone,
		Address:   user.Address,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Deleted:   false,               // Por padrão, o usuário não está deletado
		Role:      domain.RoleCustomer, // Por padrão, define o papel como cliente
	}
	if err := s.userRepo.Create(&userEntity); err != nil {
		return fmt.Errorf("error creating user: %w", err)
	}

	// Cria um carrinho para o usuário recém-criado
	if err := s.cartRepo.Create(&userEntity.ID); err != nil {
		return fmt.Errorf("error creating cart for user: %w", err)
	}

	return nil
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
		Name:    user.Name,
		Email:   user.Email,
		Role:    string(user.Role),
		Phone:   user.Phone,
		Address: user.Address,
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
		Name:    user.Name,
		Email:   user.Email,
		Role:    string(user.Role),
		Phone:   user.Phone,
		Address: user.Address,
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

	//
	if request.Email != nil {
		// Verifica se o email já está em uso por outro usuário
		existingEmailUser, err := s.userRepo.FindByEmail(*request.Email)
		if err != nil {
			return fmt.Errorf("error checking for existing user by email: %w", err)
		}
		if existingEmailUser != nil && existingEmailUser.ID != id {
			return fmt.Errorf("email already in use by another user")
		}
		existingUser.Email = *request.Email
	}

	if request.Phone != nil {
		existingUser.Phone = *request.Phone
	}
	if request.Address != nil {
		existingUser.Address = *request.Address
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

func (s *userService) Login(email, password string) (*dto.UserResponse, error) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("error finding user by email: %w", err)
	}
	if user == nil {
		return nil, fmt.Errorf("user not found")
	}

	cart, err := s.cartRepo.GetByUserID(user.ID)
	if err != nil {
		return nil, fmt.Errorf("error retrieving cart for user: %w", err)
	}
	if cart == nil {
		return nil, fmt.Errorf("cart not found for user")
	}

	if err := utils.CheckPasswordHash(&password, &user.Password); err != nil {
		return nil, fmt.Errorf("invalid password: %w", err)
	}

	return &dto.UserResponse{
		Name:    user.Name,
		Email:   user.Email,
		Role:    string(user.Role),
		Phone:   user.Phone,
		Address: user.Address,
		CartID:  cart.ID,
	}, nil
}

func (s *userService) FindUserIDByEmail(email string) (uint, error) {
	if email == "" {
		return 0, fmt.Errorf("email is required")
	}

	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return 0, fmt.Errorf("error finding user by email: %w", err)
	}
	if user == nil {
		return 0, fmt.Errorf("user not found")
	}

	return user.ID, nil
}
