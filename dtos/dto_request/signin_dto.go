package dto_request

type LoginDto struct {
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
