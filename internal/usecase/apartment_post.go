package usecase

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/kmdavidds/mager-spot-api/entity"
	"github.com/kmdavidds/mager-spot-api/internal/repository"
	"github.com/kmdavidds/mager-spot-api/model"
	"github.com/kmdavidds/mager-spot-api/pkg/supabase"
)

type IApartmentPostUsecase interface {
	CreateApartmentPost(param model.ApartmentPostCreate) error
	GetApartmentPosts() ([]entity.ApartmentPost, error)
	GetApartmentPost(key model.ApartmentPostKey) (entity.ApartmentPost, error)
	SearchApartmentPosts(key model.ApartmentPostKey) ([]entity.ApartmentPost, error)
}

type ApartmentPostUsecase struct {
	apr repository.IApartmentPostRepository
	sb  supabase.Interface
}

func NewApartmentPostUsecase(apartmentPostRepository repository.IApartmentPostRepository, supabase supabase.Interface) IApartmentPostUsecase {
	return &ApartmentPostUsecase{
		apr: apartmentPostRepository,
		sb:  supabase,
	}
}

func (apu *ApartmentPostUsecase) CreateApartmentPost(param model.ApartmentPostCreate) error {
	param.Image.Filename = fmt.Sprintf("%s.%s", strings.ReplaceAll(time.Now().String(), " ", ""), strings.Split(param.Image.Filename, ".")[1])

	imageLink, err := apu.sb.Upload(param.Image)
	if err != nil {
		return err
	}

	apartmentPost := entity.ApartmentPost{
		ID:              uuid.New(),
		UserID:          param.UserID,
		Title:           param.Title,
		Price:           param.Price,
		Body:            param.Body,
		PictureLink:     imageLink,
		Payment:         param.Payment,
		NumberOfRatings: 0,
		AverageRating:   0,
	}

	_, err = apu.apr.CreateApartmentPost(apartmentPost)
	if err != nil {
		return err
	}

	return nil
}

func (apu *ApartmentPostUsecase) GetApartmentPosts() ([]entity.ApartmentPost, error) {
	apartmentPosts, err := apu.apr.GetApartmentPosts()
	if err != nil {
		return nil, err
	}

	return apartmentPosts, nil
}

func (apu *ApartmentPostUsecase) GetApartmentPost(key model.ApartmentPostKey) (entity.ApartmentPost, error) {
	apartmentPost, err := apu.apr.GetApartmentPost(key)
	if err != nil {
		return entity.ApartmentPost{}, err
	}

	return apartmentPost, nil
}

func (apu *ApartmentPostUsecase) SearchApartmentPosts(key model.ApartmentPostKey) ([]entity.ApartmentPost, error) {
	apartmentPosts, err := apu.apr.SearchApartmentPosts(key)
	if err != nil {
		return apartmentPosts, err
	}
	return apartmentPosts, nil
}
