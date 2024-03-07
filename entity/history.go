package entity

import (
	"time"

	"github.com/google/uuid"
)

type History struct {
	ID        uuid.UUID `json:"id" gorm:"primary_key"`
	PostID    uuid.UUID `json:"postId" gorm:"primary_key;"`
	UserID    uuid.UUID `json:"userId" gorm:"primary_key;foreignkey:ID;references:users;"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
}
