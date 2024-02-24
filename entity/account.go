package entity

import (
	"github.com/google/uuid"
)

type Account struct {
	ID          uuid.UUID
	Transaction Transaction
	Balance     float32
}
