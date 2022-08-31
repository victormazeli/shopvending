package dto

type VerifyOTPDTO struct {
	OTPCode string `json:"otp_code" validate:"required|min_len:6"`
}
