package handler

import (
	"fmt"
	"net/http"
	"strconv"

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


// @Summary Find a product by ID
// @Description Retrieve a product by its ID
// @Tags Products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} dto.ProductResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /product/{id} [get]
func (h *ProductHandler) FindProductByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"error": "Product ID is required"})
		return
	}
	idInt, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": fmt.Sprintf("Invalid product ID: %v", err)})
		return
	}

	product, err := h.productService.FindProductByID(uint(idInt))
	if err != nil {
		c.JSON(500, gin.H{"error": fmt.Sprintf("Failed to find product: %v", err)})
		return
	}
	if product == nil {
		c.JSON(404, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}


// @Summary Find all products
// @Description Retrieve a list of all products
// @Tags Products
// @Accept json
// @Produce json
// @Success 200 {array} dto.ProductResponse
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /products [get]
func (h *ProductHandler) FindAllProducts(c *gin.Context) {
	products, err := h.productService.FindAllProducts()
	if err != nil {
		c.JSON(500, gin.H{"error": fmt.Sprintf("Failed to retrieve products: %v", err)})
		return 
	}
	if len(products) == 0 {
		c.JSON(404, gin.H{"error": "No products found"})
		return
	}

	c.JSON(http.StatusOK, products)
}



// @Summary Update a product
// @Description Update a product with the provided details
// @Tags Products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Param product body dto.ProductUpdateRequest true "Product details"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /product/{id} [patch]
func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"error": "Product ID is required"})
		return
	}
	
	idInt, err := strconv.ParseUint(id, 10, 32)
	var idUint uint = uint(idInt)
	if err != nil {
		c.JSON(400, gin.H{"error": fmt.Sprintf("Invalid product ID: %v", err)})
		return
	}

	var request dto.ProductUpdateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": fmt.Sprintf("Invalid request body: %v", err)})
		return
	}

	
	err = h.productService.UpdateProduct(request, &idUint)
	if err != nil {
		c.JSON(500, gin.H{"error": fmt.Sprintf("Failed to update product: %v", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully"})
}



// @Summary Delete a product
// @Description Delete a product by its ID
// @Tags Products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /product/{id} [delete]
func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"error": "Product ID is required"})
		return
	}

	idInt, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": fmt.Sprintf("Invalid product ID: %v", err)})
		return
	}

	err = h.productService.DeleteProduct(uint(idInt))
	if err != nil {
		c.JSON(500, gin.H{"error": fmt.Sprintf("Failed to delete product: %v", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}