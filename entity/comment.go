package entity

import (
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	ID        uuid.UUID `json:"id" gorm:"primary_key"`
	PostID    uuid.UUID `json:"postId" gorm:"primary_key;"`
	UserID    uuid.UUID `json:"userId" gorm:"primary_key;foreignkey:ID;references:users;"`
	Body      string    `json:"body" gorm:"not null"`
	User      User
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
}
