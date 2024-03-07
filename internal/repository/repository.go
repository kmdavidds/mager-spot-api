package repository

import "gorm.io/gorm"

type Repository struct {
	UserRepository   IUserRepository
	BarangRepository IBarangRepository
	CommentRepository ICommentRepository
}

func NewRepository(db *gorm.DB) *Repository {
	userRepository := NewUserRepository(db)
	barangRepository := NewBarangRepository(db)
	commentRepository := NewCommentRepository(db)

	return &Repository{
		UserRepository: userRepository,
		BarangRepository: barangRepository,
		CommentRepository: commentRepository,
	}
}
