package usecase

import (
	"github.com/kmdavidds/mager-spot-api/internal/repository"
	"github.com/kmdavidds/mager-spot-api/pkg/bcrypt"
	"github.com/kmdavidds/mager-spot-api/pkg/jwt_auth"
	"github.com/kmdavidds/mager-spot-api/pkg/supabase"
)

type Usecase struct {
	UserUsecase          IUserUsecase
	ApartmentPostUsecase IApartmentPostUsecase
	FoodPostUsecase      IFoodPostUsecase
	ProductPostUsecase   IProductPostUsecase
	ShuttlePostUsecase   IShuttlePostUsecase
	CommentUsecase       ICommentUsecase
	InvoiceUsecase       IInvoiceUsecase
}

type InitParam struct {
	Repository *repository.Repository
	Bcrypt     bcrypt.Interface
	JWTAuth    jwt_auth.Interface
	Supabase   supabase.Interface
}

func NewUsecase(param InitParam) *Usecase {
	userUsecase := NewUserUsecase(param.Repository.UserRepository, param.Bcrypt, param.JWTAuth, param.Supabase)
	foodPostUsecase := NewFoodPostUsecase(param.Repository.FoodPostRepository, param.Supabase)
	apartmentPostUsecase := NewApartmentPostUsecase(param.Repository.ApartmentPostRepository, param.Supabase)
	productPostUsecase := NewProductPostUsecase(param.Repository.ProductPostRepository, param.Supabase)
	shuttlePostUsecase := NewShuttlePostUsecase(param.Repository.ShuttlePostRepository, param.Supabase)
	commentUsecase := NewCommentUsecase(param.Repository.CommentRepository)
	invoiceUsecase := NewInvoiceUsecase(param.Repository.InvoiceRepository, *param.Repository)

	return &Usecase{
		UserUsecase:          userUsecase,
		ApartmentPostUsecase: apartmentPostUsecase,
		FoodPostUsecase:      foodPostUsecase,
		ProductPostUsecase:   productPostUsecase,
		ShuttlePostUsecase:   shuttlePostUsecase,
		CommentUsecase:       commentUsecase,
		InvoiceUsecase:       invoiceUsecase,
	}
}
