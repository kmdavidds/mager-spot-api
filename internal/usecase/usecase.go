package usecase

import (
	"github.com/kmdavidds/mager-spot-api/internal/repository"
	"github.com/kmdavidds/mager-spot-api/pkg/bcrypt"
	"github.com/kmdavidds/mager-spot-api/pkg/jwt_auth"
	"github.com/kmdavidds/mager-spot-api/pkg/supabase"
)

type Usecase struct {
	UserUsecase IUserUsecase
	PostUsecase IPostUsecase
}

type InitParam struct {
	Repository *repository.Repository
	Bcrypt     bcrypt.Interface
	JWTAuth    jwt_auth.Interface
	Supabase   supabase.Interface
}

func NewUsecase(param InitParam) *Usecase {
	userUsecase := NewUserUsecase(param.Repository.UserRepository, param.Bcrypt, param.JWTAuth, param.Supabase)
	postUsecase := NewPostUsecase(param.Repository.PostRepository)

	return &Usecase{
		UserUsecase: userUsecase,
		PostUsecase: postUsecase,
	}
}
