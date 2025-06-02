package dto

import "github.com/lucas-rech/sisinfo-ecommerce/internal/domain"

type ProductCreateRequest struct {
	Name        string          `json:"name" binding:"required"`
	Description string          `json:"description"`
	Price       float64         `json:"price" binding:"required,gt=0"`
	Stock       int             `json:"stock" binding:"required,gte=0"`
	ImageURL    string          `json:"image_url"`
	Category    domain.Category `json:"category" binding:"required,oneof=CLOTHING ACCESSORIES PERSONALITY"`
}

type ProductUpdateRequest struct {
	Name        *string          `json:"name"`
	Description *string          `json:"description"`
	Price       *float64         `json:"price"`
	Stock       *int             `json:"stock"`
	ImageURL    *string          `json:"image_url"`
	Category    *domain.Category `json:"category"`
}

type ProductResponse struct {
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Price       float64         `json:"price"`
	Stock       int             `json:"stock"`
	ImageURL    string          `json:"image_url"`
	Category    domain.Category `json:"category" binding:"oneof=CLOTHING ACCESSORIES PERSONALITY"`
}
