package domain

import (
	"github.com/google/uuid"
)

type Transaction struct {
	ID                   uuid.UUID `gorm:"primaryKey"`
	OriginAccountID      uuid.UUID
	OriginAccount        Account `gorm:"foreignKey:OriginAccountID;references:ID"`
	DestinationAccountID uuid.UUID
	DestinationAccount   Account `gorm:"foreignKey:DestinationAccountID;references:ID"`
	creatAt              uuid.Time
	Amount               float32
}
