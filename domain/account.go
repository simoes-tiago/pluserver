package domain

import (
	"github.com/google/uuid"
)

type Account struct {
	ID          uuid.UUID `gorm:"primaryKey"`
	Transaction Transaction
	Balance     float32
}
