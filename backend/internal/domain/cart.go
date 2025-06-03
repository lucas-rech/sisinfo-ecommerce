package domain

import "time"

type Cart struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	UserID    uint       `gorm:"uniqueIndex" json:"user_id"`
	Items     []CartItem `gorm:"foreignKey:CartID" json:"items"`
	UpdatedAt time.Time  `json:"updated_at"`
}
