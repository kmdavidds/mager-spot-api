package repository

import (
	"github.com/kmdavidds/mager-spot-api/entity"
	"github.com/kmdavidds/mager-spot-api/model"
	"gorm.io/gorm"
)

type IApartmentPostRepository interface {
	CreateApartmentPost(apartmentPost entity.ApartmentPost) (entity.ApartmentPost, error)
	GetApartmentPosts() ([]entity.ApartmentPost, error)
	GetApartmentPost(key model.ApartmentPostKey) (entity.ApartmentPost, error)
}

type ApartmentPostRepository struct {
	db *gorm.DB
}

func NewApartmentPostRepository(db *gorm.DB) IApartmentPostRepository {
	return &ApartmentPostRepository{
		db: db,
	}
}

func (apr *ApartmentPostRepository) CreateApartmentPost(apartmentPost entity.ApartmentPost) (entity.ApartmentPost, error) {
	err := apr.db.Create(&apartmentPost).Error
	if err != nil {
		return apartmentPost, err
	}

	return apartmentPost, nil
}

func (apr *ApartmentPostRepository) GetApartmentPosts() ([]entity.ApartmentPost, error) {
	apartmentPosts := []entity.ApartmentPost{}
	err := apr.db.Preload("User").Find(&apartmentPosts).Error
	if err != nil {
		return nil, err
	}
	return apartmentPosts, nil
}

func (apr *ApartmentPostRepository) GetApartmentPost(key model.ApartmentPostKey) (entity.ApartmentPost, error) {
	apartmentPost := entity.ApartmentPost{}
	err := apr.db.Preload("User").Preload("Comments").Preload("Comments.User").Where(key).First(&apartmentPost).Error
	if err != nil {
		return apartmentPost, err
	}
	return apartmentPost, nil
}
