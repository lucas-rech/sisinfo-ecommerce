package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lucas-rech/sisinfo-ecommerce/backend/internal/dto"
	"github.com/lucas-rech/sisinfo-ecommerce/backend/internal/service"
	"github.com/lucas-rech/sisinfo-ecommerce/backend/utils"
)

type UserHandler struct {
	userService service.UserService
}


func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}



// @Summary Create a new user
// @Description Create a new user with the provided details
// @Tags Users
// @Accept json
// @Produce json
// @Param user body dto.UserCreateRequest true "User details"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /user [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
	var request dto.UserCreateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	err := h.userService.CreateUser(request)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(201, gin.H{"message": "User created successfully"})
}



// @Summary Find a user by ID
// @Description Retrieve a user by their ID
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} dto.UserResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /user/{id} [get]
func (h *UserHandler) FindByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"error": "User ID is required"})
		return
	}
	idInt, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid User ID format"})
		return
	}


	user, err := h.userService.FindUserByID(uint(idInt))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, user)
}


// @Summary Find a user by email
// @Description Retrieve a user by their email address
// @Tags Users
// @Accept json
// @Produce json
// @Param email path string true "User Email"
// @Success 200 {object} dto.UserResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /user/email/{email} [get]
func (h *UserHandler) FindByEmail(c *gin.Context) {
	email := c.Param("email")
	if email == "" {
		c.JSON(400, gin.H{"error": "Email is required"})
		return
	}
	user, err := h.userService.FindUserByEmail(email)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, user)
}


// @Summary Update a user
// @Description Update user details by ID
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param user body dto.UserUpdateRequest true "User update details"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
func (h *UserHandler) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"error": "User ID is required"})
		return
	}
	idInt, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid User ID format"})
		return
	}

	var request dto.UserUpdateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	err = h.userService.UpdateUser(request, uint(idInt))
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(200, gin.H{"message": "User updated successfully"})
}


// @Summary Delete a user
// @Description Delete a user by ID
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /user/{id} [delete]
func (h *UserHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"error": "User ID is required"})
		return
	}
	idInt, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid User ID format"})
		return
	}

	err = h.userService.DeleteUser(uint(idInt))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "User deleted successfully"})
}



// @Summary User Login
// @Description Authenticate a user and return a JWT token
// @Tags Users
// @Accept json
// @Produce json
// @Param user body dto.UserLoginRequest true "User login details"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /login [post]
func (h *UserHandler) Login (c *gin.Context) {
	var request dto.UserLoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	userResponse, err := h.userService.Login(request.Email, request.Password)
	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid email or password"})
		return
	}

	token, err := utils.GenerateJWT(*userResponse)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to generate token"})
		return
	} 

	c.JSON(200, gin.H{
		"token": token,
		"user":  userResponse,
	})
}


