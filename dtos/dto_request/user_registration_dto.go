package dto_request

type UserRegistrationDto struct {
	Name string `json:"name" binding:"required,min=5,max=20"`
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}
