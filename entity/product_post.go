package entity

import (
	"time"

	"github.com/google/uuid"
)

type ProductPost struct {
	ID              uuid.UUID `json:"id" gorm:"primary_key;unique;"`
	UserID          uuid.UUID `json:"userId" gorm:"foreignkey:ID;references:users;"`
	Title           string    `json:"title" gorm:"not null"`
	Price           string    `json:"price" gorm:"not null"`
	Body            string    `json:"body" gorm:"not null"`
	PictureLink     string    `json:"pictureLink"`
	CreatedAt       time.Time `json:"createdAt" gorm:"autoCreateTime"`
	Usage           string    `json:"usage" gorm:"not null"`
	NumberOfRatings uint      `json:"numberOfRatings"`
	AverageRating   float32   `json:"averageRating"`
	User            User      `json:"user"`
	Comments        []Comment `json:"comments"`
}
