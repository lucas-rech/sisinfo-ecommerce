package dto

type CartItemCreateRequest struct {
	ProductID uint `json:"product_id" binding:"required"`
	Quantity  int  `json:"quantity" binding:"required,min=1,max=10"`
}

type CartItemUpdateRequest struct {
	ProductID uint `json:"product_id" binding:"required"`
	Quantity  int  `json:"quantity" binding:"required,max=10"`
}

type CartItemResponse struct {
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
	CartID    uint `json:"cart_id"`
}

type CartItemDeleteRequest struct {
	ProductID uint `json:"product_id" binding:"required"`
	CartID    uint `json:"cart_id" binding:"required"`
}
