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

type IMakananUsecase interface {
	GetMakanan(param model.MakananParam) (entity.MakananWithAuthor, []entity.Comment, error)
	CreateMakanan(param model.MakananCreate) error
	GetAllMakanan() ([]entity.MakananWithAuthor, error)
	ContactMakanan(param model.MakananContact) (string, error)
}

type MakananUsecase struct {
	mr repository.IMakananRepository
	sb supabase.Interface
}

func NewMakananUsecase(makananRepository repository.IMakananRepository, supabase supabase.Interface) IMakananUsecase {
	return &MakananUsecase{
		mr: makananRepository,
		sb: supabase,
	}
}

func (mu *MakananUsecase) GetMakanan(param model.MakananParam) (entity.MakananWithAuthor, []entity.Comment, error) {
	return mu.mr.GetMakanan(param)
}

func (mu *MakananUsecase) CreateMakanan(param model.MakananCreate) error {
	param.ID = uuid.New()

	param.Image.Filename = fmt.Sprintf("%s.%s", param.ID.String(), strings.Split(param.Image.Filename, ".")[1])

	imageLink, err := mu.sb.Upload(param.Image)
	if err != nil {
		return err
	}

	makanan := entity.Makanan{
		ID:          param.ID,
		UserID:      param.UserID,
		Title:       param.Title,
		Price:       param.Price,
		Body:        param.Body,
		PictureLink: imageLink,
		Varian:      param.Varian,
	}

	_, err = mu.mr.CreateMakanan(makanan)
	if err != nil {
		return err
	}

	return nil
}

func (mu *MakananUsecase) GetAllMakanan() ([]entity.MakananWithAuthor, error) {
	makanans, err := mu.mr.GetAllMakanan()
	if err != nil {
		return nil, err
	}

	return makanans, nil
}

func (mu *MakananUsecase) ContactMakanan(param model.MakananContact) (string, error) {
	makanan, user, err := mu.mr.ContactMakanan(param)
	if err != nil {
		return "", err
	}

	user.DisplayName = strings.ReplaceAll(user.DisplayName, " ", "%20")
	makanan.Title = strings.ReplaceAll(makanan.Title, " ", "%20")

	message := fmt.Sprintf("Halo%%2C%%20nama%%20saya%%20%s.%%0ASaya%%20ingin%%20tertarik%%20dengan%%20postingan%%20anda%%20yang%%20berjudul%%20%s.%%0AApakah%%20masih%%20tersedia%%3F", user.DisplayName, makanan.Title)

	contactLink := fmt.Sprintf("https://wa.me/%s?text=%s", user.PhoneNumber, message)

	return contactLink, nil
}
