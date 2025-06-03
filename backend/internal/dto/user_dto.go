package dto

type UserCreateRequest struct {
	Name     string `json:"name" binding:"required"`
	LastName string `json:"last_name" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Address  string `json:"address" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type UserUpdateRequest struct {
	Email    *string `json:"email"`
	Name     *string `json:"name"`
	Password *string `string:"password"`
	Phone    *string `json:"phone"`
	Address  *string `json:"address"`
}

type UserResponse struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	Role  string `json:"role"`
	Phone string `json:"phone"`
	Address string `json:"address"`
}

type UserLoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
