package dtos

type UserResponse struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type NewUserRequest struct {
	Email    string `json:"email" form:"email" binding:"required,email"`
	Name     string `json:"name" form:"name" binding:"required"`
	Password string `json:"password" form:"password" binding:"required,min=6"`
}
