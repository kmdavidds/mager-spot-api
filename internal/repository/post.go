package repository

import (
	"github.com/kmdavidds/mager-spot-api/entity"
	"github.com/kmdavidds/mager-spot-api/model"
	"gorm.io/gorm"
)

type IPostRepository interface {
	CreatePost(post entity.Post) (entity.Post, error)
	GetPost(param model.PostParam) (entity.Post, error)
	DeletePost(post model.PostDelete) error
}

type PostRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) IPostRepository {
	return &PostRepository{
		db: db,
	}
}

func (pr *PostRepository) CreatePost(post entity.Post) (entity.Post, error) {
	err := pr.db.Create(&post).Error
	if err != nil {
		return post, err
	}

	return post, nil
}

func (pr *PostRepository) GetPost(param model.PostParam) (entity.Post, error) {
	post := entity.Post{}
	err := pr.db.Where(&param).First(&post).Error
	if err != nil {
		return post, err
	}

	return post, nil
}

func (pr *PostRepository) DeletePost(post model.PostDelete) error {
	err := pr.db.Delete(&entity.Post{}, post.ID).Error
	if err != nil {
		return err
	}

	return nil
}
