package model

import (
	"mime/multipart"
	"time"

	"github.com/google/uuid"
)

type KosCreate struct {
	ID         uuid.UUID             `json:"-"`
	UserID     uuid.UUID             `json:"-"`
	Title      string                `form:"title" binding:"required,min=1"`
	Price      string                `form:"price" binding:"required,min=1"`
	Body       string                `form:"body" binding:"required,min=1"`
	CreatedAt  time.Time             `json:"-"`
	Pembayaran string                `form:"pembayaran" binding:"required,min=1"`
	Image      *multipart.FileHeader `form:"image"`
}

type KosParam struct {
	ID     uuid.UUID `json:"-"`
	UserID uuid.UUID `json:"-"`
}

type KosContact struct {
	ID          uuid.UUID `json:"-"`
	UserID      uuid.UUID `json:"-"`
	AskerID     uuid.UUID `json:"-"`
	ContactLink string    `json:"contactLink"`
}
