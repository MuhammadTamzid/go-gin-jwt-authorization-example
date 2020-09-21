package dto_request

type RefreshTokenRequestDto struct {
	RefreshToken string `json:"refresh_token" form:"refresh_token" binding:"required"`
}
