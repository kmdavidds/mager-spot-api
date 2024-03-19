package entity

import (
	"time"

	"github.com/google/uuid"
)

type Invoice struct {
	ID          uuid.UUID `gorm:"primaryKey"`
	UserID      uuid.UUID `gorm:"not null"`
	PostID      uuid.UUID `gorm:"not null"`
	Category    string    `gorm:"not null"`
	Status      string    `gorm:"not null"`
	PaymentLink string    `gorm:"not null"`
	CreatedAt   time.Time `gorm:"not null"`
}
