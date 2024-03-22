package model

import "github.com/google/uuid"

type InvoiceParam struct {
	ID       string
	SellerID uuid.UUID
	BuyerID  uuid.UUID
}
