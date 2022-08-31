package dto

type UpdateUserDTO struct {
	Password      string `json:"password" validate:"required|min_len:8"`
	OTPCode       string `json:"otp_code"`
	OTPExpireTime int64  `json:"otp_expire_time"`
}
