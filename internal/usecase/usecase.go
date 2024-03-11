package usecase

import (
	"github.com/kmdavidds/mager-spot-api/internal/repository"
	"github.com/kmdavidds/mager-spot-api/pkg/bcrypt"
	"github.com/kmdavidds/mager-spot-api/pkg/jwt_auth"
	"github.com/kmdavidds/mager-spot-api/pkg/supabase"
)

type Usecase struct {
	UserUsecase        IUserUsecase
	ProductPostUsecase IProductPostUsecase
	CommentUsecase     ICommentUsecase
}

type InitParam struct {
	Repository *repository.Repository
	Bcrypt     bcrypt.Interface
	JWTAuth    jwt_auth.Interface
	Supabase   supabase.Interface
}

func NewUsecase(param InitParam) *Usecase {
	userUsecase := NewUserUsecase(param.Repository.UserRepository, param.Repository.ProductPostRepository, param.Bcrypt, param.JWTAuth, param.Supabase)
	productPostUsecase := NewProductPostUsecase(param.Repository.ProductPostRepository, param.Supabase)
	commentUsecase := NewCommentUsecase(param.Repository.CommentRepository)

	return &Usecase{
		UserUsecase:        userUsecase,
		ProductPostUsecase: productPostUsecase,
		CommentUsecase:     commentUsecase,
	}
}
