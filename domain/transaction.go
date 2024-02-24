package domain

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	OriginAccountID      int64
	OriginAccount        Account `gorm:"foreignKey:OriginAccountID;references:ID"`
	DestinationAccountID int64
	DestinationAccount   Account `gorm:"foreignKey:DestinationAccountID;references:ID"`
	Amount               float32
}
