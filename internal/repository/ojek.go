package repository

import (
	"github.com/google/uuid"
	"github.com/kmdavidds/mager-spot-api/entity"
	"github.com/kmdavidds/mager-spot-api/model"
	"gorm.io/gorm"
)

type IOjekRepository interface {
	CreateOjek(ojek entity.Ojek) (entity.Ojek, error)
	GetOjek(param model.OjekParam) (entity.OjekWithAuthor, []entity.Comment, error)
	GetAllOjek() ([]entity.OjekWithAuthor, error)
	ContactOjek(param model.OjekContact) (entity.Ojek, entity.User, error)
}

type OjekRepository struct {
	db *gorm.DB
}

func NewOjekRepository(db *gorm.DB) IOjekRepository {
	return &OjekRepository{
		db: db,
	}
}

func (or *OjekRepository) CreateOjek(ojek entity.Ojek) (entity.Ojek, error) {
	err := or.db.Create(&ojek).Error
	if err != nil {
		return ojek, err
	}

	return ojek, nil
}

func (or *OjekRepository) GetOjek(param model.OjekParam) (entity.OjekWithAuthor, []entity.Comment, error) {

	ojek := entity.Ojek{}
	err := or.db.Where(&param).First(&ojek).Error
	if err != nil {
		return entity.OjekWithAuthor{}, nil, err
	}

	user := entity.User{}
	err = or.db.Where("id = ?", ojek.UserID).First(&user).Error
	if err != nil {
		return entity.OjekWithAuthor{}, nil, err
	}

	ojekWithAuthor := entity.OjekWithAuthor{
		Ojek:    ojek,
		Username:  user.Username,
		PhotoLink: user.ProfilePhotoLink,
	}

	comments := []entity.Comment{}
	err = or.db.Where("post_id = ?", ojek.ID).Find(&comments).Error
	if err != nil {
		return ojekWithAuthor, nil, err
	}

	return ojekWithAuthor, comments, nil
}

func (or *OjekRepository) GetAllOjek() ([]entity.OjekWithAuthor, error) {
	ojeks := []entity.Ojek{}
	err := or.db.Find(&ojeks).Error
	if err != nil {
		return nil, err
	}

	ojeksWithAuthor := []entity.OjekWithAuthor{}

	for _, ojek := range ojeks {
		user := entity.User{}
		err := or.db.Where("id = ?", ojek.UserID).First(&user).Error
		if err != nil {
			return nil, err
		}
		ojeksWithAuthor = append(ojeksWithAuthor, entity.OjekWithAuthor{
			Ojek:    ojek,
			Username:  user.Username,
			PhotoLink: user.ProfilePhotoLink,
		})
	}

	return ojeksWithAuthor, nil
}

func (or *OjekRepository) ContactOjek(param model.OjekContact) (entity.Ojek, entity.User, error) {
	ojek := entity.Ojek{}
	err := or.db.Where("id = ?", param.ID).First(&ojek).Error
	if err != nil {
		return entity.Ojek{}, entity.User{}, err
	}

	seller := entity.User{}
	err = or.db.Where("id = ?", ojek.UserID).First(&seller).Error
	if err != nil {
		return entity.Ojek{}, entity.User{}, err
	}

	history := entity.History{
		ID:             uuid.New(),
		PostID:         ojek.ID,
		TypeOfPost:     "ojek",
		UserID:         param.AskerID,
		Title:          ojek.Title,
		SellerUsername: seller.Username,
		Price:          ojek.Price,
	}

	err = or.db.Create(&history).Error
	if err != nil {
		return entity.Ojek{}, entity.User{}, err
	}

	return ojek, seller, nil
}
