package usecase

import (
	"github.com/google/uuid"
	"github.com/kmdavidds/mager-spot-api/entity"
	"github.com/kmdavidds/mager-spot-api/internal/repository"
	"github.com/kmdavidds/mager-spot-api/model"
)

type IBarangUsecase interface {
	GetBarang(param model.BarangParam) (entity.Barang, error)
	CreateBarang(param model.BarangCreate) error
}

type BarangUsecase struct {
	br repository.IBarangRepository
}

func NewBarangUsecase(barangRepository repository.IBarangRepository) IBarangUsecase {
	return &BarangUsecase{
		br: barangRepository,
	}
}

func (bu *BarangUsecase) GetBarang(param model.BarangParam) (entity.Barang, error) {
	return bu.br.GetBarang(param)
}

func (bu *BarangUsecase) CreateBarang(param model.BarangCreate) error {
	param.ID = uuid.New()

	post := entity.Post{
		ID:     param.ID,
		UserID: param.UserID,
		Title:  param.Title,
		Price:  param.Price,
		Body:   param.Body,
	}

	barang := entity.Barang{
		Post: post,
		Pemakaian: param.Pemakaian,
	}

	_, err := bu.br.CreateBarang(barang)
	if err != nil {
		return err
	}

	return nil
}
