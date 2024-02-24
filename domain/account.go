package domain

import (
	"github.com/google/uuid"
)

type Account struct {
	ID      uuid.UUID `gorm:"primaryKey"`
	UserID  uuid.UUID
	User    User `gorm:"foreignKey:UserID;references:ID"`
	Balance float32
}
