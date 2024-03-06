package usecase

import (
	"github.com/kmdavidds/mager-spot-api/internal/repository"
	"github.com/kmdavidds/mager-spot-api/pkg/bcrypt"
	"github.com/kmdavidds/mager-spot-api/pkg/jwt"
)

type Usecase struct {
	UserUsecase IUserUsecase
	PostUsecase IPostUsecase
}

type InitParam struct {
	Repository *repository.Repository
	Bcrypt     bcrypt.Interface
	JWT        jwt.Interface
}

func NewUsecase(param InitParam) *Usecase {
	userUsecase := NewUserUsecase(param.Repository.UserRepository, param.Bcrypt, param.JWT)
	postUsecase := NewPostUsecase(param.Repository.PostRepository)

	return &Usecase{
		UserUsecase: userUsecase,
		PostUsecase: postUsecase,
	}
}
