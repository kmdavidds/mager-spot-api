package repository

import "gorm.io/gorm"

type Repository struct {
	UserRepository IUserRepository
	PostRepository IPostRepository
}

func NewRepository(db *gorm.DB) *Repository {
	userRepository := NewUserRepository(db)
	PostRepository := NewPostRepository(db)

	return &Repository{
		UserRepository: userRepository,
		PostRepository: PostRepository,
	}
}
