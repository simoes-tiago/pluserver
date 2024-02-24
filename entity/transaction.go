package entity

import (
	"github.com/google/uuid"
)

type Transaction struct {
	ID                 uuid.UUID
	OriginAccount      uuid.UUID
	DestinationAccount uuid.UUID
	Amount             float32
}
