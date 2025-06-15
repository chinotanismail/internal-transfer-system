package models

import "time"

type Transaction struct {
	ID                   uint      `gorm:"primaryKey" json:"id"`
	SourceAccountID      uint      `json:"source_account_id"`
	DestinationAccountID uint      `json:"destination_account_id"`
	Amount               float64   `json:"amount"`
	CreatedAt            time.Time `json:"created_at"`
	Status               string    `json:"status"`
	ErrorReason          string    `json:"error_reason,omitempty"`
}
