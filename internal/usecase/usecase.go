package usecase

import (
	"github.com/kmdavidds/mager-spot-api/internal/repository"
	"github.com/kmdavidds/mager-spot-api/pkg/bcrypt"
	"github.com/kmdavidds/mager-spot-api/pkg/jwt_auth"
)

type Usecase struct {
	UserUsecase IUserUsecase
	PostUsecase IPostUsecase
}

type InitParam struct {
	Repository *repository.Repository
	Bcrypt     bcrypt.Interface
	JWTAuth    jwt_auth.Interface
}

func NewUsecase(param InitParam) *Usecase {
	userUsecase := NewUserUsecase(param.Repository.UserRepository, param.Bcrypt, param.JWTAuth)
	postUsecase := NewPostUsecase(param.Repository.PostRepository)

	return &Usecase{
		UserUsecase: userUsecase,
		PostUsecase: postUsecase,
	}
}
