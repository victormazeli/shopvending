package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type RoleAllowed string

const (
	admin   RoleAllowed = "admin"
	seller  RoleAllowed = "seller"
	rider   RoleAllowed = "rider"
	regular RoleAllowed = "regular"
)

type StatusAllowed string

const (
	suspended StatusAllowed = "suspended"
	active    StatusAllowed = "active"
	inactive  StatusAllowed = "inactive"
)

type User struct {
	ID              uint           `json:"id" gorm:"type:bigserial;primaryKey;autoIncrement"`
	Email           string         `json:"email" gorm:"type:varchar(255);unique;not null"`
	Password        string         `json:"password" gorm:"type:varchar(255);unique;not null""`
	OTPCode         int            `json:"otp_code"`
	Role            RoleAllowed    `json:"role" sql:"type:role;default:regular"`
	Status          StatusAllowed  `json:"status" sql:"type:status"`
	IsEmailVerified bool           `json:"isEmailVerified" gorm:"default:false"`
	OTPExpireDate   time.Time      `json:"otp_expire_date"`
	OTPExpireTime   int            `json:"otp_expire_time"`
	CreatedAt       time.Time      `json:"createdAt"`
	UpdatedAt       time.Time      `json:"updatedAt"`
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}

func (u *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	u.Password = string(bytes)
	return nil
}

func (u *User) CheckPasswordHash(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return err
	}
	return nil
}
