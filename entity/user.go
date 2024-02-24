package entity

import (
	"github.com/google/uuid"
)

type User struct {
	ID      uuid.UUID
	Name    string
	Country string
	Age     int
	Account Account
}
