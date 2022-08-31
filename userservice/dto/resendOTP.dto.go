package dto

type ResendOTP struct {
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}
