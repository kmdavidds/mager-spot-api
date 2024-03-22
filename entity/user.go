package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID               uuid.UUID `json:"id" gorm:"primary_key"`
	Username         string    `json:"username" gorm:"not null; unique"`
	Email            string    `json:"email" gorm:"primary_key;not null; unique" validate:"email; endswith=@student.ub.ac.id;"`
	Password         string    `json:"-" gorm:"not null"`
	DisplayName      string    `json:"displayName"`
	PhoneNumber      string    `json:"phoneNumber"`
	Address          string    `json:"Address"`
	IsSeller         bool      `json:"isSeller"`
	IsSubscribed     bool      `json:"isSubscribed"`
	SubscribedUntil  time.Time `json:"subscribedUntil"`
	ProfilePhotoLink string    `json:"profilePhotoLink"`
	Balance          int       `json:"balance"`
	Histories        []History `json:"-"`
}
