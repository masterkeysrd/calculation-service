package models

import "gorm.io/gorm"

type Operation struct {
	*gorm.Model
	OperationType string  `gorm:"not null,unique"`
	Cost          float64 `gorm:"not null,check:cost_checker,cost >= 0"`
}
