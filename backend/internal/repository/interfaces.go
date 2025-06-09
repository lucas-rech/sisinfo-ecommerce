package repository

import "github.com/lucas-rech/sisinfo-ecommerce/backend/internal/domain"

type UserRepository interface {
	Create(user *domain.User) error
	FindByID(id *uint) (*domain.User, error)
	FindByEmail(email string) (*domain.User, error)
	Update(user *domain.User) error
	Delete(id *uint) error
}

type ProductRepository interface {
	Create(product *domain.Product) error
	FindByID(id *uint) (*domain.Product, error)
	FindAll() ([]domain.Product, error)
	Update(product *domain.Product) error
	Delete(id *uint) error
}

type CartItemRepository interface {
	AddItem(item *domain.CartItem) error
	UpdateItem(item *domain.CartItem) error
	RemoveItem(cartID uint, productID uint) error
	GetItemsByCartID(cartID uint) ([]domain.CartItem, error)
	GetItemByCartAndProduct(cartID uint, productID uint) (*domain.CartItem, error)
	GetItemByUserAndProduct(userID uint, productID uint) (*domain.CartItem, error)
}

type CartRepository interface {
	Create(userId *uint) error
	GetByID(id uint) (*domain.Cart, error)
	GetByUserID(userID uint) (*domain.Cart, error)
	Update(cart *domain.Cart) error
	Delete(id uint) error
	ClearCart(cartID uint) error
}

//TODO: Checar se são apenas esses métodos necessários para implementação
type OrderItemRepository interface {
	Create(item *domain.OrderItem) error
	FindByID(id *uint) error
	FindByOrderID(orderID *uint) error
	FindByProductID(productID *uint) error
}

type OrderRepository interface {
	Create(item *domain.Order) error
	FindByID(id *uint) (*domain.Order, error)
	FindByUserID(id *uint) (*domain.Order, error)
	ChangeStatus(status *domain.OrderStatus) error
}

