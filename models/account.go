package models

type Account struct {
	AccountID uint    `gorm:"primaryKey" json:"account_id"`
	Balance   float64 `gorm:"type:numeric" json:"balance"`
}
