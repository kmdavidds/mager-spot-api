package model

import (
	"mime/multipart"
	"time"

	"github.com/google/uuid"
)

type ApartmentPostCreate struct {
	ID        uuid.UUID             `json:"-"`
	UserID    uuid.UUID             `json:"-"`
	Title     string                `form:"title" binding:"required,min=1"`
	Price     string                `form:"price" binding:"required,min=1"`
	Body      string                `form:"body" binding:"required,min=1"`
	CreatedAt time.Time             `json:"-"`
	Payment   string                `form:"payment" binding:"required,min=1"`
	Image     *multipart.FileHeader `form:"image"`
}

type ApartmentPostKey struct {
	ID       uuid.UUID
	SellerID uuid.UUID
}
