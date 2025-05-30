package domain

import "time"

type Order struct {
	ID         uint        `gorm:"primaryKey" json:"id"`
	UserID     uint        `json:"user_id"`
	TotalPrice float64     `json:"total_price"`
	Status     OrderStatus `json:"status"`
	Items      []OrdemItem `gorm:"foreignKey:OrderID" json:"items"`
	CreatedAt  time.Time   `json:"created_at"`
}
