package entity

import (
	"time"

	"github.com/google/uuid"
)

type Kos struct {
	ID          uuid.UUID `json:"id" gorm:"primary_key"`
	UserID      uuid.UUID `json:"userId" gorm:"primary_key;foreignkey:ID;references:users;"`
	Title       string    `json:"title" gorm:"not null"`
	Price       string    `json:"price" gorm:"not null"`
	Body        string    `json:"body" gorm:"not null"`
	PictureLink string    `json:"pictureLink"`
	CreatedAt   time.Time `json:"createdAt" gorm:"autoCreateTime"`
	Pembayaran  string    `json:"pembayaran" gorm:"not null"`
}

type KosWithAuthor struct {
	Kos       Kos    `json:"kos"`
	Username  string `json:"username"`
	PhotoLink string `json:"photoLink"`
}
