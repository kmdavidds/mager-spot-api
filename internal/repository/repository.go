package repository

import "gorm.io/gorm"

type Repository struct {
	UserRepository        IUserRepository
	ProductPostRepository IProductPostRepository
	CommentRepository ICommentRepository
}

func NewRepository(db *gorm.DB) *Repository {
	userRepository := NewUserRepository(db)
	productPostRepository := NewProductPostRepository(db)
	commentRepository := NewCommentRepository(db)

	return &Repository{
		UserRepository:   userRepository,
		ProductPostRepository: productPostRepository,
		CommentRepository: commentRepository,
	}
}
