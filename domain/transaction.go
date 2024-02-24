package domain

import (
	"github.com/google/uuid"
)

type Transaction struct {
	ID                 uuid.UUID
	OriginAccount      uuid.UUID
	DestinationAccount uuid.UUID
	creatAt            uuid.Time
	Amount             float32
}
