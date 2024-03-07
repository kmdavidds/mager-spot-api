package usecase

import (
	"github.com/google/uuid"
	"github.com/kmdavidds/mager-spot-api/entity"
	"github.com/kmdavidds/mager-spot-api/internal/repository"
	"github.com/kmdavidds/mager-spot-api/model"
)

type ICommentUsecase interface {
	CreateComment(param model.CommentCreate) error
	GetAllComment() ([]entity.Comment, error)
}

type CommentUsecase struct {
	cr repository.ICommentRepository
}

func NewCommentUsecase(commentRepository repository.ICommentRepository) ICommentUsecase {
	return &CommentUsecase{
		cr: commentRepository,
	}
}

func (cu *CommentUsecase) CreateComment(param model.CommentCreate) error {
	param.ID = uuid.New()

	comment := entity.Comment{
		ID: param.ID,
		PostID: param.PostID,
		UserID: param.UserID,
		Body: param.Body,
	}

	_, err := cu.cr.CreateComment(comment)
	if err != nil {
		return err
	}

	return nil
}

func (cu *CommentUsecase) GetAllComment() ([]entity.Comment, error) {
	comments, err := cu.cr.GetAllComment()
	if err != nil {
		return nil, err
	}

	return comments, nil
}
