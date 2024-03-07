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

type IOjekUsecase interface {
	GetOjek(param model.OjekParam) (entity.OjekWithAuthor, []entity.Comment, error)
	CreateOjek(param model.OjekCreate) error
	GetAllOjek() ([]entity.OjekWithAuthor, error)
	ContactOjek(param model.OjekContact) (string, error)
}

type OjekUsecase struct {
	or repository.IOjekRepository
	sb supabase.Interface
}

func NewOjekUsecase(ojekRepository repository.IOjekRepository, supabase supabase.Interface) IOjekUsecase {
	return &OjekUsecase{
		or: ojekRepository,
		sb: supabase,
	}
}

func (ou *OjekUsecase) GetOjek(param model.OjekParam) (entity.OjekWithAuthor, []entity.Comment, error) {
	return ou.or.GetOjek(param)
}

func (ou *OjekUsecase) CreateOjek(param model.OjekCreate) error {
	param.ID = uuid.New()

	param.Image.Filename = fmt.Sprintf("%s.%s", param.ID.String(), strings.Split(param.Image.Filename, ".")[1])

	imageLink, err := ou.sb.Upload(param.Image)
	if err != nil {
		return err
	}

	ojek := entity.Ojek{
		ID:          param.ID,
		UserID:      param.UserID,
		Title:       param.Title,
		Price:       param.Price,
		Body:        param.Body,
		PictureLink: imageLink,
		PlatNomor:   param.PlatNomor,
	}

	_, err = ou.or.CreateOjek(ojek)
	if err != nil {
		return err
	}

	return nil
}

func (ou *OjekUsecase) GetAllOjek() ([]entity.OjekWithAuthor, error) {
	ojeks, err := ou.or.GetAllOjek()
	if err != nil {
		return nil, err
	}

	return ojeks, nil
}

func (ou *OjekUsecase) ContactOjek(param model.OjekContact) (string, error) {
	ojek, user, err := ou.or.ContactOjek(param)
	if err != nil {
		return "", err
	}

	user.DisplayName = strings.ReplaceAll(user.DisplayName, " ", "%20")
	ojek.Title = strings.ReplaceAll(ojek.Title, " ", "%20")

	message := fmt.Sprintf("Halo%%2C%%20nama%%20saya%%20%s.%%0ASaya%%20ingin%%20tertarik%%20dengan%%20postingan%%20anda%%20yang%%20berjudul%%20%s.%%0AApakah%%20masih%%20tersedia%%3F", user.DisplayName, ojek.Title)

	contactLink := fmt.Sprintf("https://wa.me/%s?text=%s", user.PhoneNumber, message)

	return contactLink, nil
}
