package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucas-rech/sisinfo-ecommerce/internal/dto"
	"github.com/lucas-rech/sisinfo-ecommerce/internal/service"
)

type ProductHandler struct {
	productService service.ProductService
}

func NewProductHandler(productService service.ProductService) *ProductHandler {
	return &ProductHandler{productService: productService}
}

// @Summary Create a new product
// @Description Create a new product with the provided details
// @Tags Products
// @Accept json
// @Produce json
// @Param product body dto.ProductCreateRequest true "Product details"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /product [post]
func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var request dto.ProductCreateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": fmt.Sprintf("Invalid request body: %v", err)})
		return
	}

	err := h.productService.CreateProduct(request)
	if err != nil {
		c.JSON(500, gin.H{"error": fmt.Sprintf("Failed to create product: %v", err)})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Product created successfully"})
}