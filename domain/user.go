package domain

import (
	"github.com/google/uuid"
)

type User struct {
	ID      uuid.UUID `gorm:"primaryKey"`
	Phone   string    `gorm:"unique"`
	Name    string
	Country string
	Age     int
}
