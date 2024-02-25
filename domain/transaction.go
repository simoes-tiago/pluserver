package domain

import (
	"gorm.io/gorm"
)

type TransactionType int

const (
	Transfer TransactionType = iota
	Deposit
	Withdraw
)

type Transaction struct {
	gorm.Model
	ID                   uint    `gorm:"primaryKey;autoincrement"`
	OriginAccountID      *int64  `json:"origin-account-id,omitempty"`
	OriginAccount        Account `json:",omitempty" gorm:"foreignKey:OriginAccountID;references:ID"`
	DestinationAccountID *int64  `json:"destination-account-id"`
	DestinationAccount   Account `gorm:"foreignKey:DestinationAccountID;references:ID"`
	Amount               float32
	Type                 TransactionType
}
