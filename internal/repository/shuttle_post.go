package repository

import (
	"github.com/kmdavidds/mager-spot-api/entity"
	"github.com/kmdavidds/mager-spot-api/model"
	"gorm.io/gorm"
)

type IShuttlePostRepository interface {
	CreateShuttlePost(shutllePost entity.ShuttlePost) (entity.ShuttlePost, error)
	GetShuttlePosts() ([]entity.ShuttlePost, error)
	GetShuttlePost(key model.ShuttlePostKey) (entity.ShuttlePost, error)
}

type ShuttlePostRepository struct {
	db *gorm.DB
}

func NewShuttlePostRepository(db *gorm.DB) IShuttlePostRepository {
	return &ShuttlePostRepository{
		db: db,
	}
}

func (spr *ShuttlePostRepository) CreateShuttlePost(shutllePost entity.ShuttlePost) (entity.ShuttlePost, error) {
	err := spr.db.Create(&shutllePost).Error
	if err != nil {
		return shutllePost, err
	}

	return shutllePost, nil
}

func (spr *ShuttlePostRepository) GetShuttlePosts() ([]entity.ShuttlePost, error) {
	shutllePosts := []entity.ShuttlePost{}
	err := spr.db.Preload("User").Find(&shutllePosts).Error
	if err != nil {
		return nil, err
	}
	return shutllePosts, nil
}

func (spr *ShuttlePostRepository) GetShuttlePost(key model.ShuttlePostKey) (entity.ShuttlePost, error) {
	shutllePost := entity.ShuttlePost{}
	err := spr.db.Preload("User").Preload("Comments").Preload("Comments.User").Where(key).First(&shutllePost).Error
	if err != nil {
		return shutllePost, err
	}
	return shutllePost, nil
}
