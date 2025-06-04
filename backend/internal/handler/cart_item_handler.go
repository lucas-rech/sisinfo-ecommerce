package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lucas-rech/sisinfo-ecommerce/backend/internal/dto"
	"github.com/lucas-rech/sisinfo-ecommerce/backend/internal/service"
)

type CartItemHandler struct {
	cartItemService service.CartItemService
	userService     service.UserService
}

func NewCartItemHandler(cartItemService service.CartItemService, userService service.UserService) *CartItemHandler {
	return &CartItemHandler{
		cartItemService: cartItemService,
		userService:     userService,
	}
}

// @Summary Add an item to the user's cart
// @Description Add a product to the user's cart with a specified quantity
// @Tags Cart Items
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.CartItemCreateRequest true "Cart Item Create Request"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /cart/item [post]
func (h *CartItemHandler) AddItemToCart(c *gin.Context) {
	userSignature, exists := c.Get("email")
	if !exists || userSignature == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	email := userSignature.(string)
	userID, err := h.userService.FindUserIDByEmail(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user ID", "details": err.Error()})
		return
	}

	var request dto.CartItemCreateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "details": err.Error()})
		return
	}

	if err := h.cartItemService.AddItemToCart(userID, request.ProductID, request.Quantity); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add item to cart", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Item added to cart successfully"})
}

// @Summary Update an item in the user's cart
// @Description Update the quantity of a product in the user's cart
// @Tags Cart Items
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.CartItemUpdateRequest true "Cart Item Update Request"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /cart/item [patch]
func (h *CartItemHandler) UpdateItemInCart(c *gin.Context) {
	userSignature, exists := c.Get("email")
	if !exists || userSignature == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	email := userSignature.(string)
	userID, err := h.userService.FindUserIDByEmail(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user ID", "details": err.Error()})
		return
	}

	var request dto.CartItemUpdateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "details": err.Error()})
		return
	}

	if err = h.cartItemService.UpdateItemInCart(request, userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update item in cart", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item updated in cart successfully"})
}

// @Summary Remove an item from the user's cart
// @Description Remove a product from user's cart. It uses the product ID to identify which item to remove and assumes the user is authenticated with a valid JWT token.
// @Tags Cart Items
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param product_id query string true "Product ID to remove from cart"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /cart/item/{item_id} [delete]
func (h *CartItemHandler) RemoveItemFromCart(c *gin.Context) {
	cartIDValue, exists := c.Get("cart_ID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	productIDParam := c.Query("product_id")
	if productIDParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product ID is required"})
		return
	}

	productID, err := strconv.ParseUint(productIDParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID format", "details": err.Error()})
		return
	}

	requestDTO := dto.CartItemDeleteRequest{
		CartID:    cartIDValue.(uint),
		ProductID: uint(productID),
	}

	if err := h.cartItemService.RemoveItemFromCart(&requestDTO); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove item from cart", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item removed from cart successfully"})
}

// @Summary Get all items in the user's cart
// @Description Retrieve all items in the user's cart
// @Tags Cart Items
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {array} dto.CartItemResponse
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /cart/items [get]
func (h *CartItemHandler) GetAllItemsInCart(c *gin.Context) {
	cartID, exists := c.Get("cart_ID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	items, err := h.cartItemService.GetItemsByCartID(cartID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve cart items", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, items)
}
