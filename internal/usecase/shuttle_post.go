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

type IShuttlePostUsecase interface {
	CreateShuttlePost(param model.ShuttlePostCreate) error
	GetShuttlePosts() ([]entity.ShuttlePost, error)
	GetShuttlePost(key model.ShuttlePostKey) (entity.ShuttlePost, error)
}

type ShuttlePostUsecase struct {
	spr repository.IShuttlePostRepository
	sb  supabase.Interface
}

func NewShuttlePostUsecase(shuttlePostRepository repository.IShuttlePostRepository, supabase supabase.Interface) IShuttlePostUsecase {
	return &ShuttlePostUsecase{
		spr: shuttlePostRepository,
		sb:  supabase,
	}
}

func (spu *ShuttlePostUsecase) CreateShuttlePost(param model.ShuttlePostCreate) error {
	param.Image.Filename = fmt.Sprintf("%s.%s", strings.ReplaceAll(time.Now().String(), " ", ""), strings.Split(param.Image.Filename, ".")[1])

	imageLink, err := spu.sb.Upload(param.Image)
	if err != nil {
		return err
	}

	shuttlePost := entity.ShuttlePost{
		ID:              uuid.New(),
		UserID:          param.UserID,
		Title:           param.Title,
		Price:           param.Price,
		Body:            param.Body,
		PictureLink:     imageLink,
		VehicleNumber:   param.VehicleNumber,
		NumberOfRatings: 0,
		AverageRating:   0,
	}

	_, err = spu.spr.CreateShuttlePost(shuttlePost)
	if err != nil {
		return err
	}

	return nil
}

func (spu *ShuttlePostUsecase) GetShuttlePosts() ([]entity.ShuttlePost, error) {
	shuttlePosts, err := spu.spr.GetShuttlePosts()
	if err != nil {
		return nil, err
	}

	return shuttlePosts, nil
}

func (spu *ShuttlePostUsecase) GetShuttlePost(key model.ShuttlePostKey) (entity.ShuttlePost, error) {
	shuttlePost, err := spu.spr.GetShuttlePost(key)
	if err != nil {
		return entity.ShuttlePost{}, err
	}

	return shuttlePost, nil
}
