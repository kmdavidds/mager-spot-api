package repository

import (
	"github.com/kmdavidds/mager-spot-api/entity"
	"github.com/kmdavidds/mager-spot-api/model"
	"gorm.io/gorm"
)

type IUserRepository interface {
	CreateUser(user entity.User) (entity.User, error)
	GetUser(param model.UserParam) (entity.User, error)
	UpdateUser(param model.UserUpdates, user entity.User) error
	UpdatePhoto(param model.PhotoUpdate) error
	ShowHistory(user entity.User) ([]entity.History, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) CreateUser(user entity.User) (entity.User, error) {
	err := ur.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (ur *UserRepository) GetUser(param model.UserParam) (entity.User, error) {
	user := entity.User{}
	err := ur.db.Where(&param).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (ur *UserRepository) UpdateUser(param model.UserUpdates, user entity.User) error {
	err := ur.db.Model(&user).Updates(param).Error
	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) UpdatePhoto(param model.PhotoUpdate) error {
	err := ur.db.Model(&entity.User{}).Where("id = ?", param.UserID).Update("profile_photo_link", param.PhotoLink).Error
	if err != nil {
		return err
	}
	return nil
}

func (ur *UserRepository) ShowHistory(user entity.User) ([]entity.History, error) {
	histories := []entity.History{}
	err := ur.db.Where("user_id = ?", user.ID).Find(&histories).Error
	if err != nil {
		return nil, err
	}

	return histories, nil
}
