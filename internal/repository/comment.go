package repository

import (
	"github.com/kmdavidds/mager-spot-api/entity"
	"gorm.io/gorm"
)

type ICommentRepository interface {
	CreateComment(comment entity.Comment) (entity.Comment, error)
	GetAllComment() ([]entity.Comment, error)
}

type CommentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) ICommentRepository {
	return &CommentRepository{
		db: db,
	}
}

func (cr *CommentRepository) CreateComment(comment entity.Comment) (entity.Comment, error) {
	err := cr.db.Create(&comment).Error
	if err != nil {
		return comment, err
	}

	return comment, nil
}

func (cr *CommentRepository) GetAllComment() ([]entity.Comment, error) {
	comments := []entity.Comment{}
	err := cr.db.Find(&comments).Error
	if err != nil {
		return comments, err
	}
	return comments, nil
}
