package repository

import (
	"github.com/kmdavidds/mager-spot-api/entity"
	"github.com/kmdavidds/mager-spot-api/model"
	"gorm.io/gorm"
)

type IFoodPostRepository interface {
	CreateFoodPost(foodPost entity.FoodPost) (entity.FoodPost, error)
	GetFoodPosts() ([]entity.FoodPost, error)
	GetFoodPost(key model.FoodPostKey) (entity.FoodPost, error)
}

type FoodPostRepository struct {
	db *gorm.DB
}

func NewFoodPostRepository(db *gorm.DB) IFoodPostRepository {
	return &FoodPostRepository{
		db: db,
	}
}

func (fpr *FoodPostRepository) CreateFoodPost(foodPost entity.FoodPost) (entity.FoodPost, error) {
	err := fpr.db.Create(&foodPost).Error
	if err != nil {
		return foodPost, err
	}

	return foodPost, nil
}

func (fpr *FoodPostRepository) GetFoodPosts() ([]entity.FoodPost, error) {
	foodPosts := []entity.FoodPost{}
	err := fpr.db.Preload("User").Find(&foodPosts).Error
	if err != nil {
		return nil, err
	}
	return foodPosts, nil
}

func (fpr *FoodPostRepository) GetFoodPost(key model.FoodPostKey) (entity.FoodPost, error) {
	foodPost := entity.FoodPost{}
	err := fpr.db.Preload("User").Preload("Comments").Preload("Comments.User").Where(key).First(&foodPost).Error
	if err != nil {
		return foodPost, err
	}
	return foodPost, nil
}
