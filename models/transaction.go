package models

import "time"

type Transaction struct {
	ID                   uint `gorm:"primaryKey"`
	SourceAccountID      uint
	DestinationAccountID uint
	Amount               float64
	CreatedAt            time.Time `json:"created_at"`
	Status               string    `json:"status"`
	ErrorReason          string    `json:"error_reason,omitempty"`
}
