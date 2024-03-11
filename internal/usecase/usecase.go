package usecase

import (
	"github.com/kmdavidds/mager-spot-api/internal/repository"
	"github.com/kmdavidds/mager-spot-api/pkg/bcrypt"
	"github.com/kmdavidds/mager-spot-api/pkg/jwt_auth"
	"github.com/kmdavidds/mager-spot-api/pkg/supabase"
)

type Usecase struct {
	UserUsecase    IUserUsecase
	BarangUsecase  IBarangUsecase
	// KosUsecase     IKosUsecase
	// MakananUsecase IMakananUsecase
	// OjekUsecase    IOjekUsecase
	CommentUsecase ICommentUsecase
}

type InitParam struct {
	Repository *repository.Repository
	Bcrypt     bcrypt.Interface
	JWTAuth    jwt_auth.Interface
	Supabase   supabase.Interface
}

func NewUsecase(param InitParam) *Usecase {
	userUsecase := NewUserUsecase(param.Repository.UserRepository, param.Bcrypt, param.JWTAuth, param.Supabase)
	barangUsecase := NewBarangUsecase(param.Repository.BarangRepository, param.Supabase)
	// kosUsecase := NewKosUsecase(param.Repository.KosRepository, param.Supabase)
	// makananUsecase := NewMakananUsecase(param.Repository.MakananRepository, param.Supabase)
	// ojekUsecase := NewOjekUsecase(param.Repository.OjekRepository, param.Supabase)
	commentUsecase := NewCommentUsecase(param.Repository.CommentRepository)

	return &Usecase{
		UserUsecase:    userUsecase,
		BarangUsecase:  barangUsecase,
		// KosUsecase:     kosUsecase,
		// MakananUsecase: makananUsecase,
		// OjekUsecase:    ojekUsecase,
		CommentUsecase: commentUsecase,
	}
}
