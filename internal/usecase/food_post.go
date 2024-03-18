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

type IFoodPostUsecase interface {
	CreateFoodPost(param model.FoodPostCreate) error
	GetFoodPosts() ([]entity.FoodPost, error)
	GetFoodPost(key model.FoodPostKey) (entity.FoodPost, error)
	SearchFoodPosts(key model.FoodPostKey) ([]entity.FoodPost, error)
}

type FoodPostUsecase struct {
	fpr repository.IFoodPostRepository
	sb  supabase.Interface
}

func NewFoodPostUsecase(foodPostRepository repository.IFoodPostRepository, supabase supabase.Interface) IFoodPostUsecase {
	return &FoodPostUsecase{
		fpr: foodPostRepository,
		sb:  supabase,
	}
}

func (fpu *FoodPostUsecase) CreateFoodPost(param model.FoodPostCreate) error {
	param.Image.Filename = fmt.Sprintf("%s.%s", strings.ReplaceAll(time.Now().String(), " ", ""), strings.Split(param.Image.Filename, ".")[1])

	imageLink, err := fpu.sb.Upload(param.Image)
	if err != nil {
		return err
	}

	foodPost := entity.FoodPost{
		ID:              uuid.New(),
		UserID:          param.UserID,
		Title:           param.Title,
		Price:           param.Price,
		Body:            param.Body,
		PictureLink:     imageLink,
		Variant:         param.Variant,
		NumberOfRatings: 0,
		AverageRating:   0,
	}

	_, err = fpu.fpr.CreateFoodPost(foodPost)
	if err != nil {
		return err
	}

	return nil
}

func (fpu *FoodPostUsecase) GetFoodPosts() ([]entity.FoodPost, error) {
	foodPosts, err := fpu.fpr.GetFoodPosts()
	if err != nil {
		return nil, err
	}

	return foodPosts, nil
}

func (fpu *FoodPostUsecase) GetFoodPost(key model.FoodPostKey) (entity.FoodPost, error) {
	foodPost, err := fpu.fpr.GetFoodPost(key)
	if err != nil {
		return entity.FoodPost{}, err
	}

	return foodPost, nil
}

func (fpu *FoodPostUsecase) SearchFoodPosts(key model.FoodPostKey) ([]entity.FoodPost, error) {
	foodPosts, err := fpu.fpr.SearchFoodPosts(key)
	if err != nil {
		return foodPosts, err
	}
	return foodPosts, nil
}