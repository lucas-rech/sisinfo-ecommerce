package repository

import (
	"github.com/lucas-rech/sisinfo-ecommerce/backend/internal/domain"
	"gorm.io/gorm"
)

type orderItemRepository struct {
	db *gorm.DB
}

func NewOrderItemRepository(db *gorm.DB) OrderItemRepository {
	return &orderItemRepository{db: db}
}

// Create implements OrderItemRepository.
func (o *orderItemRepository) Create(item *domain.OrderItem) error {
	panic("unimplemented")
}

// FindByID implements OrderItemRepository.
func (o *orderItemRepository) FindByID(id *uint) error {
	panic("unimplemented")
}

// FindByOrderID implements OrderItemRepository.
func (o *orderItemRepository) FindByOrderID(orderID *uint) error {
	panic("unimplemented")
}

// FindByProductID implements OrderItemRepository.
func (o *orderItemRepository) FindByProductID(productID *uint) error {
	panic("unimplemented")
}
