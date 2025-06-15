package models

import "time"

type Account struct {
	AccountID uint      `gorm:"primaryKey" json:"account_id"`
	Balance   float64   `gorm:"type:numeric" json:"balance"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
