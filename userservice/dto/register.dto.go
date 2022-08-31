package dto

import "time"

type CreateNewUserDTO struct {
	Email         string    `json:"email" validate:"required|email"`
	Password      string    `json:"password" validate:"required|min_len:8"`
	OTPCode       string    `json:"otp_code"`
	OTPExpireTime time.Time `json:"otp_expire_time"`
}
