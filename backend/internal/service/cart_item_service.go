package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/lucas-rech/sisinfo-ecommerce/backend/internal/domain"
	"github.com/lucas-rech/sisinfo-ecommerce/backend/internal/dto"
	"github.com/lucas-rech/sisinfo-ecommerce/backend/internal/repository"
	"gorm.io/gorm"
)

type cartItemService struct {
	cartItemRepo repository.CartItemRepository
	cartRepo     repository.CartRepository
	productService ProductService
}

func NewCartItemService(cartItemRepo repository.CartItemRepository, cartRepo repository.CartRepository, productService ProductService) CartItemService {
	return &cartItemService{
		cartItemRepo: cartItemRepo,
		cartRepo:     cartRepo,
		productService: productService,
	}
}

func (s *cartItemService) AddItemToCart(userID uint, productID uint, quantity int) error {
	if userID == 0 || productID == 0 || quantity <= 0 {
		return fmt.Errorf("invalid user ID, product ID or quantity")
	}

	// Busca (ou cria) carrinho
	cart, err := s.cartRepo.GetByUserID(userID)
	if err != nil {
		return fmt.Errorf("failed to retrieve cart: %w", err)
	}

	// Busca item
	item, err := s.cartItemRepo.GetItemByCartAndProduct(cart.ID, productID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	// Atualiza o item se já existir
	if item != nil {
		item.Quantity += quantity
		return s.cartItemRepo.UpdateItem(item)
	}

	newItem := &domain.CartItem{
		CartID:    cart.ID,
		ProductID: productID,
		Quantity:  quantity,
		DateAdded: time.Now(),
	}

	if err := s.productService.DecreaseProductStock(productID, quantity); err != nil {
		return fmt.Errorf("failed to decrease product stock: %w", err)
	}

	return s.cartItemRepo.AddItem(newItem)
}

func (s *cartItemService) RemoveItemFromCart(request *dto.CartItemDeleteRequest) error {

	if request == nil || request.CartID == 0 || request.ProductID == 0 {
		return fmt.Errorf("invalid cart ID or product ID")
	}

	item, err := s.cartItemRepo.GetItemByCartAndProduct(request.CartID, request.ProductID)
	if err != nil {
		return fmt.Errorf("failed to retrieve item: %w", err)
	}
	//Adiciona o estoque do produto novamente
	if err := s.productService.IncreaseProductStock(request.ProductID, item.Quantity); err != nil {
		return fmt.Errorf("failed to increase product stock: %w", err)
	}

	return s.cartItemRepo.RemoveItem(request.CartID, request.ProductID)
}

func (s *cartItemService) GetItemsByCartID(cartID uint) ([]dto.CartItemResponse, error) {
	if cartID == 0 {
		return nil, fmt.Errorf("invalid cart ID")
	}

	items, err := s.cartItemRepo.GetItemsByCartID(cartID)
	if err != nil {
		return nil, err
	}

	dtoItems := make([]dto.CartItemResponse, len(items))
	for i, item := range items {
		dtoItems[i] = dto.CartItemResponse{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			CartID:    item.CartID,
		}
	}

	return dtoItems, nil
}

func (s *cartItemService) UpdateItemInCart(request dto.CartItemUpdateRequest, userID uint) error {
	if userID == 0 || request.ProductID == 0 {
		return fmt.Errorf("invalid cart ID, or product ID")
	}

	item, err := s.cartItemRepo.GetItemByUserAndProduct(userID, request.ProductID)
	if err != nil {
		return fmt.Errorf("failed to retrieve item: %w", err)
	}

	if item == nil {
		return fmt.Errorf("item not found in cart")
	}

	if request.Quantity <= 0 {
		err := s.productService.IncreaseProductStock(request.ProductID, item.Quantity)
		if err != nil {
			return fmt.Errorf("failed to increase product stock: %w", err)
		}
	} else {
		err := s.productService.DecreaseProductStock(request.ProductID, request.Quantity)
		if err != nil {
			return fmt.Errorf("failed to decrease product stock: %w", err)
		}
	}

	item.Quantity = item.Quantity + request.Quantity

	// Se a quantidade for 0, remove o item
	if item.Quantity <= 0 {
		return s.cartItemRepo.RemoveItem(item.CartID, item.ProductID)
	}

	return s.cartItemRepo.UpdateItem(item)
}
