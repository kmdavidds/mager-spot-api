package model

import (
	"mime/multipart"
	"time"

	"github.com/google/uuid"
)

type BarangCreate struct {
	ID        uuid.UUID             `json:"-"`
	UserID    uuid.UUID             `json:"-"`
	Title     string                `form:"title" binding:"required,min=1"`
	Price     string                `form:"price" binding:"required,min=1"`
	Body      string                `form:"body" binding:"required,min=1"`
	CreatedAt time.Time             `json:"-"`
	Pemakaian string                `form:"pemakaian" binding:"required,min=1"`
	Image     *multipart.FileHeader `form:"image"`
}

type BarangParam struct {
	ID     uuid.UUID `json:"-"`
	UserID uuid.UUID `json:"-"`
}

type BarangContact struct {
	ID          uuid.UUID `json:"-"`
	UserID      uuid.UUID `json:"-"`
	ContactLink string    `json:"contactLink"`
}
