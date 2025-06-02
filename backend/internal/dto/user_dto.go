package dto

type UserCreateRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type UserUpdateRequest struct {
	Email    *string `json:"email"`
	Name     *string `json:"name"`
	Password *string `string:"password"`
}

type UserResponse struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	Role  string `json:"role"`
}
