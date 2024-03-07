package entity

import (
	"time"

	"github.com/google/uuid"
)

type History struct {
	ID             uuid.UUID `json:"id" gorm:"primary_key"`
	PostID         uuid.UUID `json:"postId" gorm:"primary_key;"`
	TypeOfPost     string    `json:"typeOfPost"`
	UserID         uuid.UUID `json:"userId" gorm:"primary_key;foreignkey:ID;references:users;"`
	CreatedAt      time.Time `json:"createdAt" gorm:"autoCreateTime"`
	Title          string    `json:"title"`
	SellerUsername string    `json:"sellerUsername"`
	Price          string    `json:"price"`
}

type HistoryWithSeller struct {
	History        History `json:"history"`
	Title          string  `json:"title"`
	SellerUsername string  `json:"sellerUsername"`
	Price          string  `json:"price"`
}

type HistoryAll struct {
	HistoryBarang  []HistoryWithSeller
	HistoryKos     []HistoryWithSeller
	HistoryMakanan []HistoryWithSeller
	HistoryOjek    []HistoryWithSeller
}
