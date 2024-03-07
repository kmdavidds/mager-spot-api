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

type IKosUsecase interface {
	GetKos(param model.KosParam) (entity.KosWithAuthor, []entity.Comment, error)
	CreateKos(param model.KosCreate) error
	GetAllKos() ([]entity.KosWithAuthor, error)
	ContactKos(param model.KosContact) (string, error)
}

type KosUsecase struct {
	kr repository.IKosRepository
	sb supabase.Interface
}

func NewKosUsecase(kosRepository repository.IKosRepository, supabase supabase.Interface) IKosUsecase {
	return &KosUsecase{
		kr: kosRepository,
		sb: supabase,
	}
}

func (ku *KosUsecase) GetKos(param model.KosParam) (entity.KosWithAuthor, []entity.Comment, error) {
	return ku.kr.GetKos(param)
}

func (ku *KosUsecase) CreateKos(param model.KosCreate) error {
	param.ID = uuid.New()

	param.Image.Filename = fmt.Sprintf("%s.%s", param.ID.String(), strings.Split(param.Image.Filename, ".")[1])

	imageLink, err := ku.sb.Upload(param.Image)
	if err != nil {
		return err
	}

	kos := entity.Kos{
		ID:          param.ID,
		UserID:      param.UserID,
		Title:       param.Title,
		Price:       param.Price,
		Body:        param.Body,
		PictureLink: imageLink,
		Pembayaran:  param.Pembayaran,
	}

	_, err = ku.kr.CreateKos(kos)
	if err != nil {
		return err
	}

	return nil
}

func (ku *KosUsecase) GetAllKos() ([]entity.KosWithAuthor, error) {
	koss, err := ku.kr.GetAllKos()
	if err != nil {
		return nil, err
	}

	return koss, nil
}

func (ku *KosUsecase) ContactKos(param model.KosContact) (string, error) {
	kos, user, err := ku.kr.ContactKos(param)
	if err != nil {
		return "", err
	}

	user.DisplayName = strings.ReplaceAll(user.DisplayName, " ", "%20")
	kos.Title = strings.ReplaceAll(kos.Title, " ", "%20")

	message := fmt.Sprintf("Halo%%2C%%20nama%%20saya%%20%s.%%0ASaya%%20ingin%%20tertarik%%20dengan%%20postingan%%20anda%%20yang%%20berjudul%%20%s.%%0AApakah%%20masih%%20tersedia%%3F", user.DisplayName, kos.Title)

	contactLink := fmt.Sprintf("https://wa.me/%s?text=%s", user.PhoneNumber, message)

	return contactLink, nil
}
