package repository

import "gorm.io/gorm"

type Repository struct {
	UserRepository   IUserRepository
	PostRepository   IPostRepository
	BarangRepository IBarangRepository
}

func NewRepository(db *gorm.DB) *Repository {
	userRepository := NewUserRepository(db)
	postRepository := NewPostRepository(db)
	barangRepository := NewBarangRepository(db)

	return &Repository{
		UserRepository: userRepository,
		PostRepository: postRepository,
		BarangRepository: barangRepository,
	}
}
