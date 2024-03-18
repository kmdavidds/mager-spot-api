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

type IProductPostUsecase interface {
	CreateProductPost(param model.ProductPostCreate) error
	GetProductPosts() ([]entity.ProductPost, error)
	GetProductPost(key model.ProductPostKey) (entity.ProductPost, error)
	SearchProductPosts(key model.ProductPostKey) ([]entity.ProductPost, error)
}

type ProductPostUsecase struct {
	ppr repository.IProductPostRepository
	sb  supabase.Interface
}

func NewProductPostUsecase(productPostRepository repository.IProductPostRepository, supabase supabase.Interface) IProductPostUsecase {
	return &ProductPostUsecase{
		ppr: productPostRepository,
		sb:  supabase,
	}
}

func (ppu *ProductPostUsecase) CreateProductPost(param model.ProductPostCreate) error {
	param.Image.Filename = fmt.Sprintf("%s.%s", strings.ReplaceAll(time.Now().String(), " ", ""), strings.Split(param.Image.Filename, ".")[1])

	imageLink, err := ppu.sb.Upload(param.Image)
	if err != nil {
		return err
	}

	productPost := entity.ProductPost{
		ID:              uuid.New(),
		UserID:          param.UserID,
		Title:           param.Title,
		Price:           param.Price,
		Body:            param.Body,
		PictureLink:     imageLink,
		Usage:           param.Usage,
		NumberOfRatings: 0,
		AverageRating:   0,
	}

	_, err = ppu.ppr.CreateProductPost(productPost)
	if err != nil {
		return err
	}

	return nil
}

func (ppu *ProductPostUsecase) GetProductPosts() ([]entity.ProductPost, error) {
	productPosts, err := ppu.ppr.GetProductPosts()
	if err != nil {
		return nil, err
	}

	return productPosts, nil
}

func (ppu *ProductPostUsecase) GetProductPost(key model.ProductPostKey) (entity.ProductPost, error) {
	productPost, err := ppu.ppr.GetProductPost(key)
	if err != nil {
		return entity.ProductPost{}, err
	}

	return productPost, nil
}

func (ppu *ProductPostUsecase) SearchProductPosts(key model.ProductPostKey) ([]entity.ProductPost, error) {
	foodPosts, err := ppu.ppr.SearchProductPosts(key)
	if err != nil {
		return foodPosts, err
	}
	return foodPosts, nil
}