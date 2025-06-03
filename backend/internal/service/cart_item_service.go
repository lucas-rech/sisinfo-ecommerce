package service

import (
	"fmt"

	"github.com/lucas-rech/sisinfo-ecommerce/backend/internal/domain"
	"github.com/lucas-rech/sisinfo-ecommerce/backend/internal/repository"
)


type CartItemService struct {
	cartItemRepo repository.CartItemRepository
}

func NewCartItemService(cartItemRepo repository.CartItemRepository) *CartItemService {
	return &CartItemService{
		cartItemRepo: cartItemRepo,
	}
}

func (s *CartItemService) AddItemToCart(cartID uint, productID uint, quantity int) error {
	if cartID == 0 || productID == 0 || quantity <= 0 {
		return fmt.Errorf("invalid cart ID, product ID or quantity")
	}

	// Busca pelo item no carrinho
	item, err := s.cartItemRepo.GetItemByCartAndProduct(cartID, productID)
	if err != nil {
		return err
	}

	// Se o item já existe, apenas atualiza a quantidade
	if item != nil {
		item.Quantity += quantity
		return s.cartItemRepo.AddItem(item)
	}

	newItem := &domain.CartItem{
		CartID:    cartID,
		ProductID: productID,
		Quantity:  quantity,
	}
	return s.cartItemRepo.AddItem(newItem)
}

func (s *CartItemService) RemoveItemFromCart(cartID uint, productID uint) error {
	if cartID == 0 || productID == 0 {
		return fmt.Errorf("invalid cart ID or product ID")
	}

	return s.cartItemRepo.RemoveItem(cartID, productID)
}


func (s *CartItemService) GetItemsByCartID(cartID uint) ([]domain.CartItem, error) {
	if cartID == 0 {
		return nil, fmt.Errorf("invalid cart ID")
	}

	items, err := s.cartItemRepo.GetItemsByCartID(cartID)
	if err != nil {
		return nil, err
	}

	return items, nil
}