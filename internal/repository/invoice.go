package repository

import (
	"github.com/kmdavidds/mager-spot-api/entity"
	"gorm.io/gorm"
)

type IInvoiceRepository interface {
	CreateInvoice(invoice entity.Invoice) error
	UpdateInvoiceStatus(status string, id string) error
}

type InvoiceRepository struct {
	db *gorm.DB
}

func NewInvoiceRepository(db *gorm.DB) IInvoiceRepository {
	return &InvoiceRepository{
		db: db,
	}
}

func (ir *InvoiceRepository) CreateInvoice(invoice entity.Invoice) error {
	err := ir.db.Create(&invoice).Error
	if err != nil {
		return err
	}

	return nil
}

func (ir *InvoiceRepository) UpdateInvoiceStatus(status string, id string) error {
	err := ir.db.Model(&entity.Invoice{}).Where("id = ?", id).Update("status", status).Error
	if err != nil {
		return err
	}

	return nil
}
