package repository

import (
	"fmt"
	"time"

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

	items, err := c.GetItemsByCartID(item.CartID)
	if err != nil {
		return fmt.Errorf("error retrieving items for cart: %w", err)
	}

	// Checa se o item já existe no carrinho, se existir, atualiza a quantidade e a data de atualização
	for _, existingItem := range items {
		if existingItem.ProductID == item.ProductID {
			existingItem.Quantity += item.Quantity
			existingItem.DateAdded = time.Now()
			return c.UpdateItem(&existingItem)
		}
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

func (c *cartItemRepository) GetItemByUserAndProduct(userID uint, productID uint) (*domain.CartItem, error) {
	if userID == 0 && productID == 0 {
		return nil, fmt.Errorf("invalid user or product ID")
	}

	var item domain.CartItem
	err := c.db.
	Joins("JOIN carts ON carts.id = cart_items.cart_id").
	Where("carts.user_id = ? AND cart_items.product_id = ?", userID, productID).
	First(&item).Error
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



func (c *cartItemRepository) UpdateItem(item *domain.CartItem) error {
	if item == nil || item.ID == 0 {
		return fmt.Errorf("invalid cart item")
	}

	result := c.db.Model(&domain.CartItem{}).
		Where("id = ?", item.ID).
		Updates(map[string]interface{}{
			"quantity":   item.Quantity,
			"date_added": time.Now(),
		})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("no item found with the given ID")
	}
	return nil
}


