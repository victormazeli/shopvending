package dto

type ForgotPasswordDTO struct {
	Email string `json:"email" validate:"required|email"`
}
