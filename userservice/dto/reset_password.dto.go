package dto

type ResetPasswordDTO struct {
	NewPasssword string `json:"new_passsword" validate:"required|min_len:8"`
	OTPCode      string `json:"otp_code" validate:"required"`
}
