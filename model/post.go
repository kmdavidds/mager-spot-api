package model

import (
	"mime/multipart"
	"time"

	"github.com/google/uuid"
)

type PostCreate struct {
	ID         uuid.UUID             `json:"-"`
	UserID     uuid.UUID             `json:"-"`
	Title      string                `json:"title" binding:"required,min=1"`
	Price      string                `json:"price" binding:"required,min=1"`
	Body       string                `json:"body" binding:"required,min=1"`
	Photo      *multipart.FileHeader `form:"photo"`
	CreatedAt  time.Time             `json:"-"`
	Details    string                `json:"details" binding:"required,min=1"`
	TypeOfPost string                `json:"typeOfPost" binding:"required,min=1"`
}

type PostParam struct {
	ID     uuid.UUID `json:"-"`
	UserID uuid.UUID `json:"-"`
}

type PostDelete struct {
	ID        uuid.UUID `json:"id" binding:"required,min=1"`
	UserID    uuid.UUID `json:"-"`
	Title     string    `json:"-"`
	Body      string    `json:"-"`
	CreatedAt time.Time `json:"-"`
}

type UserUploadPhoto struct {
	Photo *multipart.FileHeader `form:"photo"`
}
