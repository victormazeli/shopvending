package dto

type CreateNewUserDTO struct {
	Email    string `json:"email" validate:"required|email"`
	Password string `json:"password" validate:"required|min_len:8"`
}
