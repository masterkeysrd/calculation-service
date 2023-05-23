package models

import "gorm.io/gorm"

type Record struct {
	*gorm.Model
	UserID      uint      `gorm:"not null"`
	User        User      `gorm:"foreignKey:UserID"`
	OperationID uint      `gorm:"not null"`
	Operation   Operation `gorm:"foreignKey:OperationID"`
	Amount      float64   `gorm:"not null"`
	UserBalance float64   `gorm:"not null"`
	Result      string    `gorm:"not null"`
}
