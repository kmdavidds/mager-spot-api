package repository

import (
	"github.com/kmdavidds/mager-spot-api/entity"
	"github.com/kmdavidds/mager-spot-api/model"
	"gorm.io/gorm"
)

type IBarangRepository interface {
	CreateBarang(barang entity.Barang) (entity.Barang, error)
	GetBarang(param model.BarangParam) (entity.Barang, error)
	GetAllBarang() ([]entity.Barang, error)
}

type BarangRepository struct {
	db *gorm.DB
}

func NewBarangRepository(db *gorm.DB) IBarangRepository {
	return &BarangRepository{
		db: db,
	}
}

func (br *BarangRepository) CreateBarang(barang entity.Barang) (entity.Barang, error) {
	err := br.db.Create(&barang).Error
	if err != nil {
		return barang, err
	}

	return barang, nil
}

func (br *BarangRepository) GetBarang(param model.BarangParam) (entity.Barang, error) {
	barang := entity.Barang{}
	err := br.db.Where(&param).First(&barang).Error
	if err != nil {
		return barang, err
	}

	return barang, nil
}

func (br *BarangRepository) GetAllBarang() ([]entity.Barang, error) {
	barangs := []entity.Barang{}
	err := br.db.Find(&barangs).Error
	if err != nil {
		return barangs, err
	}
	return barangs, nil
}
