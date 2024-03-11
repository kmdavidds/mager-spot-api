package model

import (
	"mime/multipart"
	"time"

	"github.com/google/uuid"
)

type ShuttlePostCreate struct {
	ID            uuid.UUID             `json:"-"`
	UserID        uuid.UUID             `json:"-"`
	Title         string                `form:"title" binding:"required,min=1"`
	Price         string                `form:"price" binding:"required,min=1"`
	Body          string                `form:"body" binding:"required,min=1"`
	CreatedAt     time.Time             `json:"-"`
	VehicleNumber string                `form:"vehicleNumber" binding:"required,min=1"`
	Image         *multipart.FileHeader `form:"image"`
}

type ShuttlePostKey struct {
	ID       uuid.UUID
	SellerID uuid.UUID
}
