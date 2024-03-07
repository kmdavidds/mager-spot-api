package repository

import "gorm.io/gorm"

type Repository struct {
	UserRepository    IUserRepository
	BarangRepository  IBarangRepository
	KosRepository     IKosRepository
	CommentRepository ICommentRepository
}

func NewRepository(db *gorm.DB) *Repository {
	userRepository := NewUserRepository(db)
	barangRepository := NewBarangRepository(db)
	kosRepository := NewKosRepository(db)
	commentRepository := NewCommentRepository(db)

	return &Repository{
		UserRepository:    userRepository,
		BarangRepository:  barangRepository,
		KosRepository:     kosRepository,
		CommentRepository: commentRepository,
	}
}
