package model

import (
	"time"

	"github.com/google/uuid"
)

type CommentCreate struct {
	ID              uuid.UUID `json:"id"`
	ApartmentPostID uuid.UUID `json:"apartmentPostID"`
	FoodPostID      uuid.UUID `json:"foodPostID"`
	ProductPostID   uuid.UUID `json:"productPostID"`
	ShuttlePostID   uuid.UUID `json:"shuttlePostID"`
	UserID          uuid.UUID `json:"userID"`
	Body            string    `json:"body"`
	CreatedAt       time.Time `json:"-"`
}

type CommentParam struct {
	ID     uuid.UUID `json:"-"`
	PostID uuid.UUID `json:"-"`
}
