package dto

type RefreshTokenDTO struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type ValidateRefreshToken struct {
	UserId    int  `json:"user_id"`
	IsRefresh bool `json:"user_type"`
}
