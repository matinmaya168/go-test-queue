package models

import (
	"time"
)

// Payment represents a payment in the queue
// swagger:model Payment
type Payment struct {
	// ID of the payment
	// example: 1
	ID uint `json:"id" gorm:"primaryKey"`
	// User ID associated with the payment
	// example: user1
	// required: true
	UserID string `json:"user_id" gorm:"not null" validate:"required,alphanum"`
	// Amount of the payment
	// example: 99.99
	// required: true
	// minimum: 0
	Amount float64 `json:"amount" gorm:"not null" validate:"required,gt=0"`
	// Product ID associated with the payment
	// example: item1
	// required: true
	ProductID string `json:"product_id" gorm:"not null" validate:"required,alphanum"`
	// Status of the payment
	// example: pending
	// enum: pending,processed,failed
	// required: true
	Status string `json:"status" gorm:"default:pending" validate:"oneof=pending processed failed"`
	// Timestamp when the payment was created
	// example: 2025-09-09T17:04:00Z
	// format: date-time
	Timestamp time.Time `json:"timestamp" gorm:"autoCreateTime"`
	// Creation time (from gorm.Model)
	// example: 2025-09-09T17:04:00Z
	// format: date-time
	CreatedAt time.Time `json:"created_at"`
	// Update time (from gorm.Model)
	// example: 2025-09-09T17:04:00Z
	// format: date-time
	UpdatedAt time.Time `json:"updated_at"`
	// Soft delete time (nullable)
	// example: null
	// format: date-time
	DeletedAt *time.Time `json:"deleted_at" gorm:"index"`
}
