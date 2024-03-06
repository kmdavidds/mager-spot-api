package usecase

import (
	"github.com/kmdavidds/mager-spot-api/entity"
	"github.com/kmdavidds/mager-spot-api/internal/repository"
	"github.com/kmdavidds/mager-spot-api/model"
)

type IPostUsecase interface {
	GetPost(param model.PostParam) (entity.Post, error)
	CreatePost(param model.PostCreate) error
	DeletePost(param model.PostDelete) error
}

type PostUsecase struct {
	pr repository.IPostRepository
}

func NewPostUsecase(postRepository repository.IPostRepository) IPostUsecase {
	return &PostUsecase{
		pr: postRepository,
	}
}

func (pu *PostUsecase) GetPost(param model.PostParam) (entity.Post, error) {
	return pu.pr.GetPost(param)
}

func (pu *PostUsecase) CreatePost(param model.PostCreate) error {
	return nil
}

func (pu *PostUsecase) DeletePost(param model.PostDelete) error {
	return pu.pr.DeletePost(param)
}
