package models

type Transaction struct {
	ID                   uint `gorm:"primaryKey"`
	SourceAccountID      uint
	DestinationAccountID uint
	Amount               float64
}
