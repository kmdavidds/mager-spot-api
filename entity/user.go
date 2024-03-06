package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID               uuid.UUID `json:"id" gorm:"primary_key"`
	Username         string    `json:"username" gorm:"not null; unique"`
	Email            string    `json:"email" gorm:"not null; unique" validate:"email; endswith=@student.ub.ac.id;"`
	Password         string    `json:"password" gorm:"not null"`
	DisplayName      string    `json:"-"`
	PhoneNumber      string    `json:"-"`
	Address          string    `json:"-"`
	IsSeller         bool      `json:"-"`
	IsSubscribed     bool      `json:"-"`
	SubscribedUntil  time.Time `json:"-"`
	ProfilePhotoLink string    `json:"-"`
}
