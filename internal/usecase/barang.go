package usecase

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/kmdavidds/mager-spot-api/entity"
	"github.com/kmdavidds/mager-spot-api/internal/repository"
	"github.com/kmdavidds/mager-spot-api/model"
	"github.com/kmdavidds/mager-spot-api/pkg/supabase"
)

type IBarangUsecase interface {
	GetBarang(param model.BarangParam) (entity.Barang, error)
	CreateBarang(param model.BarangCreate) error
	GetAllBarang() ([]entity.Barang, error)
}

type BarangUsecase struct {
	br repository.IBarangRepository
	sb supabase.Interface
}

func NewBarangUsecase(barangRepository repository.IBarangRepository, supabase supabase.Interface) IBarangUsecase {
	return &BarangUsecase{
		br: barangRepository,
		sb: supabase,
	}
}

func (bu *BarangUsecase) GetBarang(param model.BarangParam) (entity.Barang, error) {
	return bu.br.GetBarang(param)
}

func (bu *BarangUsecase) CreateBarang(param model.BarangCreate) error {
	param.ID = uuid.New()

	param.Image.Filename = fmt.Sprintf("%s.%s", param.ID.String(), strings.Split(param.Image.Filename, ".")[1])

	imageLink, err := bu.sb.Upload(param.Image)
	if err != nil {
		return err
	}

	barang := entity.Barang{
		ID:          param.ID,
		UserID:      param.UserID,
		Title:       param.Title,
		Price:       param.Price,
		Body:        param.Body,
		PictureLink: imageLink,
		Pemakaian: param.Pemakaian,
	}

	_, err = bu.br.CreateBarang(barang)
	if err != nil {
		return err
	}

	return nil
}

func (bu *BarangUsecase) GetAllBarang() ([]entity.Barang, error) {
	barangs, err := bu.br.GetAllBarang()
	if err != nil {
		return nil, err
	}

	return barangs, nil
}
