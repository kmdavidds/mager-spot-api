package model

import (
	"time"

	"github.com/google/uuid"
)

type BarangCreate struct {
	ID        uuid.UUID `json:"-"`
	UserID    uuid.UUID `json:"-"`
	Title     string    `json:"title" binding:"required,min=1"`
	Price     string    `json:"price" binding:"required,min=1"`
	Body      string    `json:"body" binding:"required,min=1"`
	CreatedAt time.Time `json:"-"`
	Pemakaian string    `json:"pemakaian" binding:"required,min=1"`
}

type BarangParam struct {
	ID     uuid.UUID `json:"-"`
	UserID uuid.UUID `json:"-"`
}
