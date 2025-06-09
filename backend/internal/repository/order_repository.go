package repository

import (
	"fmt"

	"github.com/lucas-rech/sisinfo-ecommerce/backend/internal/domain"
	"gorm.io/gorm"
)

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db: db}
}

// ChangeStatus implements OrderRepository.
func (o *orderRepository) ChangeStatus(status *domain.OrderStatus) error {
	panic("unimplemented")
}

func (o *orderRepository) Create(order *domain.Order) error {
	if order == nil {
		return fmt.Errorf("product cannot be null")
	}
	// Gera exceção caso o id do usuário não exista
	var user domain.User
	if err := o.db.Where("user_id = ?", order.UserID).First(&user).Error; err != nil {
		return fmt.Errorf("user not found")
	}

	return o.db.Create(order).Error
}

func (o *orderRepository) FindByID(id *uint) (*domain.Order, error) {
	if id == nil {
		return nil, fmt.Errorf("invalid order ID")
	}

	var order domain.Order
	err := o.db.First(&order, id).Error
	if err != nil {
		return nil, err
	}

	return &order, nil
}

// FindByUserID implements OrderRepository.
func (o *orderRepository) FindByUserID(id *uint) (*domain.Order, error) {
	panic("unimplemented")
}
