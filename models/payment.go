package models

import (
	"time"

	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	UserID    string    `json:"user_id" gorm:"not null" validate:"required,alphanum"`
	Amount    float64   `json:"amount" gorm:"not null" validate:"required,gt=0"`
	ProductID string    `json:"product_id" gorm:"not null" validate:"required,alphanum"`
	Status    string    `json:"status" gorm:"default:pending" validate:"oneof=pending processed failed"`
	Timestamp time.Time `json:"timestamp" gorm:"autoCreateTime"`
}
