package repository

import (
	"github.com/google/uuid"
	"github.com/kmdavidds/mager-spot-api/entity"
	"github.com/kmdavidds/mager-spot-api/model"
	"gorm.io/gorm"
)

type IKosRepository interface {
	CreateKos(Kos entity.Kos) (entity.Kos, error)
	GetKos(param model.KosParam) (entity.KosWithAuthor, []entity.Comment, error)
	GetAllKos() ([]entity.KosWithAuthor, error)
	ContactKos(param model.KosContact) (entity.Kos, entity.User, error)
}

type KosRepository struct {
	db *gorm.DB
}

func NewKosRepository(db *gorm.DB) IKosRepository {
	return &KosRepository{
		db: db,
	}
}

func (kr *KosRepository) CreateKos(kos entity.Kos) (entity.Kos, error) {
	err := kr.db.Create(&kos).Error
	if err != nil {
		return kos, err
	}

	return kos, nil
}

func (kr *KosRepository) GetKos(param model.KosParam) (entity.KosWithAuthor, []entity.Comment, error) {

	kos := entity.Kos{}
	err := kr.db.Where(&param).First(&kos).Error
	if err != nil {
		return entity.KosWithAuthor{}, nil, err
	}

	user := entity.User{}
	err = kr.db.Where("id = ?", kos.UserID).First(&user).Error
	if err != nil {
		return entity.KosWithAuthor{}, nil, err
	}

	barangWithAuthor := entity.KosWithAuthor{
		Kos:    kos,
		Username:  user.Username,
		PhotoLink: user.ProfilePhotoLink,
	}

	comments := []entity.Comment{}
	err = kr.db.Where("post_id = ?", kos.ID).Find(&comments).Error
	if err != nil {
		return barangWithAuthor, nil, err
	}

	return barangWithAuthor, comments, nil
}

func (kr *KosRepository) GetAllKos() ([]entity.KosWithAuthor, error) {
	barangs := []entity.Kos{}
	err := kr.db.Find(&barangs).Error
	if err != nil {
		return nil, err
	}

	barangsWithAuthor := []entity.KosWithAuthor{}

	for _, kos := range barangs {
		user := entity.User{}
		err := kr.db.Where("id = ?", kos.UserID).First(&user).Error
		if err != nil {
			return nil, err
		}
		barangsWithAuthor = append(barangsWithAuthor, entity.KosWithAuthor{
			Kos:    kos,
			Username:  user.Username,
			PhotoLink: user.ProfilePhotoLink,
		})
	}

	return barangsWithAuthor, nil
}

func (kr *KosRepository) ContactKos(param model.KosContact) (entity.Kos, entity.User, error) {
	kos := entity.Kos{}
	err := kr.db.Where("id = ?", param.ID).First(&kos).Error
	if err != nil {
		return entity.Kos{}, entity.User{}, err
	}

	seller := entity.User{}
	err = kr.db.Where("id = ?", kos.UserID).First(&seller).Error
	if err != nil {
		return entity.Kos{}, entity.User{}, err
	}

	history := entity.History{
		ID:             uuid.New(),
		PostID:         kos.ID,
		TypeOfPost:     "kos",
		UserID:         param.AskerID,
		Title:          kos.Title,
		SellerUsername: seller.Username,
		Price:          kos.Price,
	}

	err = kr.db.Create(&history).Error
	if err != nil {
		return entity.Kos{}, entity.User{}, err
	}

	return kos, seller, nil
}
