package entity

import (
	"time"

	"github.com/google/uuid"
)

type Invoice struct {
	ID            uuid.UUID `gorm:"primaryKey"`
	UserID        uuid.UUID `gorm:"not null"`
	SellerID      uuid.UUID
	PostID        uuid.UUID `gorm:"not null"`
	Category      string    `gorm:"not null"`
	OriginalPrice int64
	Status        string `gorm:"not null"`
	Amount        string
	PaymentLink   string    `gorm:"not null"`
	CreatedAt     time.Time `gorm:"not null"`
}
