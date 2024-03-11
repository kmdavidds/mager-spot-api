package entity

import (
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	ID              uuid.UUID `json:"id" gorm:"primary_key"`
	ApartmentPostID uuid.UUID `json:"apartmentPostId" gorm:"foreignkey:ID;references:apartment_posts;"`
	FoodPostID      uuid.UUID `json:"foodPostId" gorm:"foreignkey:ID;references:food_posts;"`
	ProductPostID   uuid.UUID `json:"productPostId" gorm:"foreignkey:ID;references:product_posts;"`
	ShuttlePostID   uuid.UUID `json:"shuttlePostId" gorm:"foreignkey:ID;references:product_posts;"`
	UserID          uuid.UUID `json:"userId" gorm:"primary_key;foreignkey:ID;references:users;"`
	Body            string    `json:"body" gorm:"not null"`
	User            User      `json:"user"`
	CreatedAt       time.Time `json:"createdAt" gorm:"autoCreateTime"`
}
