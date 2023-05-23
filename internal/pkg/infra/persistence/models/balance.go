package models

import "gorm.io/gorm"

type Balance struct {
	*gorm.Model
	UserID         uint    `gorm:"not null"`
	Amount         float64 `gorm:"not null"`
	AmountInFlight float64 `gorm:"not null"`
}
