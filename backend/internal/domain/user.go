package domain

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	LastName  string    `json:"last_name"`
	Email     string    `gorm:"unique" json:"email"`
	Password  string    `json:"-"`
	Role      Role      `json:"role"`
	Cart      *Cart     `json:"cart,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Address   string    `json:"address,omitempty"`
	Phone     string    `json:"phone,omitempty"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	Deleted   bool      `json:"deleted,omitempty" gorm:"default:false"`
}
