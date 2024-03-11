package repository

import (
	"github.com/kmdavidds/mager-spot-api/entity"
	"github.com/kmdavidds/mager-spot-api/model"
	"gorm.io/gorm"
)

type IProductPostRepository interface {
	CreateProductPost(productPost entity.ProductPost) (entity.ProductPost, error)
	GetProductPosts() ([]entity.ProductPost, error)
	GetProductPost(key model.ProductPostKey) (entity.ProductPost, error)
}

type ProductPostRepository struct {
	db *gorm.DB
}

func NewProductPostRepository(db *gorm.DB) IProductPostRepository {
	return &ProductPostRepository{
		db: db,
	}
}

func (ppr *ProductPostRepository) CreateProductPost(productPost entity.ProductPost) (entity.ProductPost, error) {
	err := ppr.db.Create(&productPost).Error
	if err != nil {
		return productPost, err
	}

	return productPost, nil
}

func (ppr *ProductPostRepository) GetProductPosts() ([]entity.ProductPost, error) {
	productPosts := []entity.ProductPost{}
	err := ppr.db.Preload("User").Find(&productPosts).Error
	if err != nil {
		return nil, err
	}
	return productPosts, nil
}

func (ppr *ProductPostRepository) GetProductPost(key model.ProductPostKey) (entity.ProductPost, error) {
	productPost := entity.ProductPost{}
	err := ppr.db.Preload("User").Preload("Comments").Preload("Comments.User").Where(key).First(&productPost).Error
	if err != nil {
		return productPost, err
	}
	return productPost, nil
}
