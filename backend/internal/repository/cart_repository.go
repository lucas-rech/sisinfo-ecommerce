package repository

import (
	"github.com/lucas-rech/sisinfo-ecommerce/backend/internal/domain"
	"gorm.io/gorm"
)

type cartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) CartRepository {
	return &cartRepository{db: db}
}


func (c *cartRepository) ClearCart(cartID uint) error {
	if cartID == 0 {
		return gorm.ErrInvalidData
	}

	return c.db.Where("cart_id = ?", cartID).Delete(&domain.CartItem{}).Error
}

func (c *cartRepository) Create(userID *uint) error {
	if userID == nil {
		return gorm.ErrInvalidData
	}
	cart := &domain.Cart{
		UserID: *userID,
		Items:  []domain.CartItem{},
	}
	
	if err := c.db.Where("user_id = ?", userID).First(cart).Error; err == nil {
		// Carrinho já existe para o usuário, não cria um novo
		return nil
	}

	return c.db.Create(cart).Error
}

func (c *cartRepository) Delete(id uint) error {
	if id == 0 {
		return gorm.ErrInvalidData
	}

	return c.db.Delete(&domain.Cart{}, id).Error
}


func (c *cartRepository) GetByID(id uint) (*domain.Cart, error) {
	if id == 0 {
		return nil, gorm.ErrInvalidData
	}

	var cart domain.Cart
	err := c.db.Preload("Items").First(&cart, id).Error
	if err != nil {
		return nil, err
	}
	return &cart, nil
}

func (c *cartRepository) GetByUserID(userID uint) (*domain.Cart, error) {
	if userID == 0 {
		return nil, gorm.ErrInvalidData
	}

	var cart domain.Cart
	err := c.db.Preload("Items").Where("user_id = ?", userID).First(&cart).Error
	if err != nil {
		return nil, err
	}
	return &cart, nil
}

// Update implements CartRepository.
func (c *cartRepository) Update(cart *domain.Cart) error {
	if cart == nil || cart.ID == 0 {
		return gorm.ErrInvalidData
	}

	return c.db.Save(cart).Error
}

