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
	GetBarang(param model.BarangParam) (entity.BarangWithAuthor, []entity.Comment, error)
	CreateBarang(param model.BarangCreate) error
	GetAllBarang() ([]entity.BarangWithAuthor, error)
	ContactBarang(param model.BarangContact) (string, error)
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

func (bu *BarangUsecase) GetBarang(param model.BarangParam) (entity.BarangWithAuthor, []entity.Comment, error) {
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
		Pemakaian:   param.Pemakaian,
	}

	_, err = bu.br.CreateBarang(barang)
	if err != nil {
		return err
	}

	return nil
}

func (bu *BarangUsecase) GetAllBarang() ([]entity.BarangWithAuthor, error) {
	barangs, err := bu.br.GetAllBarang()
	if err != nil {
		return nil, err
	}

	return barangs, nil
}

func (bu *BarangUsecase) ContactBarang(param model.BarangContact) (string, error) {
	barang, user, err := bu.br.ContactBarang(param)
	if err != nil {
		return "", err
	}

	user.DisplayName = strings.ReplaceAll(user.DisplayName, " ", "%20")
	barang.Title = strings.ReplaceAll(barang.Title, " ", "%20")

	message := fmt.Sprintf("Halo%%2C%%20nama%%20saya%%20%s.%%0ASaya%%20ingin%%20tertarik%%20dengan%%20postingan%%20anda%%20yang%%20berjudul%%20%s.%%0AApakah%%20masih%%20tersedia%%3F", user.DisplayName, barang.Title)

	contactLink := fmt.Sprintf("https://wa.me/%s?text=%s", user.PhoneNumber, message)

	return contactLink, nil
}
