package repository

import (
	"github.com/kmdavidds/mager-spot-api/entity"
	"github.com/kmdavidds/mager-spot-api/model"
	"gorm.io/gorm"
)

type IInvoiceRepository interface {
	CreateInvoice(invoice entity.Invoice) error
	UpdateInvoiceStatus(status string, id string) (*gorm.DB, error)
	GetInvoice(param model.InvoiceParam) (entity.Invoice, error)
	AddBalance(tx *gorm.DB, invoice entity.Invoice) error
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

func (ir *InvoiceRepository) UpdateInvoiceStatus(status string, id string) (*gorm.DB, error) {
	tx := ir.db.Begin()
	tx = tx.Model(&entity.Invoice{}).Where("id = ?", id).Update("status", status)
	if tx.Error != nil {
		return tx, tx.Error
	}

	return tx, tx.Error
}

func (ir *InvoiceRepository) GetInvoice(param model.InvoiceParam) (entity.Invoice, error) {
	invoice := entity.Invoice{}
	err := ir.db.Where(&param).First(&invoice).Error
	if err != nil {
		return invoice, err
	}

	return invoice, nil
}
func (ir *InvoiceRepository) AddBalance(tx *gorm.DB, invoice entity.Invoice) error {
	err := tx.Model(&entity.User{}).Where("id = ?", invoice.SellerID).Update("balance", gorm.Expr("balance + ?", invoice.OriginalPrice)).Error
	if err != nil {
		return err
	}
	return nil
}
