package domain

import "time"

type CartItem struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CartID    uint      `json:"cart_id"`
	ProductID uint      `json:"product_id"`
	Quantity  int       `json:"quantity"`
	Product   Product   `gorm:"foreignKey:ProductID" json:"product"`
	DateAdded time.Time `gorm:"autoCreateTime" json:"date_added"`
}
