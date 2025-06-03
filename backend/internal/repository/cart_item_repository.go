package repository

import (
	"fmt"

	"github.com/lucas-rech/sisinfo-ecommerce/backend/internal/domain"
	"gorm.io/gorm"
)

type cartItemRepository struct {
	db *gorm.DB
}

func NewCartItemRepository(db *gorm.DB) CartItemRepository {
	return &cartItemRepository{db: db}
}

func (c *cartItemRepository) AddItem(item *domain.CartItem) error {
	if item == nil {
		return fmt.Errorf("product cannot be nil")
	}

	return c.db.Create(item).Error
}

func (c *cartItemRepository) GetItemByCartAndProduct(cartID uint, productID uint) (*domain.CartItem, error) {
	if cartID == 0 || productID == 0 {
		return nil, fmt.Errorf("invalid cart ID or product ID")
	}

	var item domain.CartItem
	err := c.db.Where("cart_id = ? AND product_id = ?", cartID, productID).First(&item).Error
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (c *cartItemRepository) GetItemsByCartID(cartID uint) ([]domain.CartItem, error) {
	if cartID == 0 {
		return nil, fmt.Errorf("invalid cart ID")
	}

	var items []domain.CartItem
	err := c.db.Where("cart_id = ?", cartID).Find(&items).Error
	if err != nil {
		return nil, err
	}
	return items, nil
}

// RemoveItem implements CartItemRepository.
func (c *cartItemRepository) RemoveItem(cartID uint, productID uint) error {
	if cartID == 0 || productID == 0 {
		return fmt.Errorf("invalid cart ID or product ID")
	}

	result := c.db.Where("cart_id = ? AND product_id = ?", cartID, productID).Delete(&domain.CartItem{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("no item found with the given cart ID and product ID")
	}
	return nil
}

// UpdateItem implements CartItemRepository.
func (c *cartItemRepository) UpdateItem(item *domain.CartItem) error {
	if item == nil || item.ID == 0 {
		return fmt.Errorf("invalid cart item")
	}

	result := c.db.Save(item)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("no item found with the given ID")
	}
	return nil
}


