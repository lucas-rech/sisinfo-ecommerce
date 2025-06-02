package repository

import (
	"fmt"

	"github.com/lucas-rech/sisinfo-ecommerce/internal/domain"
	"gorm.io/gorm"
)

//Análogo a uma classe privada
type userRepository struct {
	db *gorm.DB
}

//Análogo a uma interface
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}



func (r *userRepository) Create(user *domain.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) Delete(id *uint) error {
	if id == nil {
		return fmt.Errorf("invalid user ID for deletion")
	}

	return r.db.Delete(&domain.User{}, id).Error
}

func (r *userRepository) FindByEmail(email string) (*domain.User, error) {
	var user domain.User
	err := r.db.Where("email =?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindByID(id *uint) (*domain.User, error) {
	if id == nil {
		return nil, fmt.Errorf("invalid user ID")
	}
	
	var user domain.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Update(user *domain.User) error {
	if user == nil || user.ID == 0 {
		return fmt.Errorf("user to update cannot be nil and must have a valid ID")
	}

	result := r.db.Save(user)
	if result.Error != nil {
		return fmt.Errorf("error updating user: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no user found with ID %d to update", user.ID)
	}

	return nil
}




