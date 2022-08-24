package common

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID        int `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
