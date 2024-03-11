package repository

import "gorm.io/gorm"

type Repository struct {
	UserRepository          IUserRepository
	ApartmentPostRepository IApartmentPostRepository
	FoodPostRepository      IFoodPostRepository
	ProductPostRepository   IProductPostRepository
	ShuttlePostRepository   IShuttlePostRepository
	CommentRepository       ICommentRepository
}

func NewRepository(db *gorm.DB) *Repository {
	userRepository := NewUserRepository(db)
	apartmentPostRepository := NewApartmentPostRepository(db)
	foodPostRepository := NewFoodPostRepository(db)
	productPostRepository := NewProductPostRepository(db)
	shuttlePostRepository := NewShuttlePostRepository(db)
	commentRepository := NewCommentRepository(db)

	return &Repository{
		UserRepository:          userRepository,
		ApartmentPostRepository: apartmentPostRepository,
		FoodPostRepository:      foodPostRepository,
		ProductPostRepository:   productPostRepository,
		ShuttlePostRepository:   shuttlePostRepository,
		CommentRepository:       commentRepository,
	}
}
