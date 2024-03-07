package repository

import (
	"github.com/google/uuid"
	"github.com/kmdavidds/mager-spot-api/entity"
	"github.com/kmdavidds/mager-spot-api/model"
	"gorm.io/gorm"
)

type IMakananRepository interface {
	CreateMakanan(makanan entity.Makanan) (entity.Makanan, error)
	GetMakanan(param model.MakananParam) (entity.MakananWithAuthor, []entity.Comment, error)
	GetAllMakanan() ([]entity.MakananWithAuthor, error)
	ContactMakanan(param model.MakananContact) (entity.Makanan, entity.User, error)
}

type MakananRepository struct {
	db *gorm.DB
}

func NewMakananRepository(db *gorm.DB) IMakananRepository {
	return &MakananRepository{
		db: db,
	}
}

func (mr *MakananRepository) CreateMakanan(makanan entity.Makanan) (entity.Makanan, error) {
	err := mr.db.Create(&makanan).Error
	if err != nil {
		return makanan, err
	}

	return makanan, nil
}

func (mr *MakananRepository) GetMakanan(param model.MakananParam) (entity.MakananWithAuthor, []entity.Comment, error) {

	makanan := entity.Makanan{}
	err := mr.db.Where(&param).First(&makanan).Error
	if err != nil {
		return entity.MakananWithAuthor{}, nil, err
	}

	user := entity.User{}
	err = mr.db.Where("id = ?", makanan.UserID).First(&user).Error
	if err != nil {
		return entity.MakananWithAuthor{}, nil, err
	}

	makananWithAuthor := entity.MakananWithAuthor{
		Makanan:    makanan,
		Username:  user.Username,
		PhotoLink: user.ProfilePhotoLink,
	}

	comments := []entity.Comment{}
	err = mr.db.Where("post_id = ?", makanan.ID).Find(&comments).Error
	if err != nil {
		return makananWithAuthor, nil, err
	}

	return makananWithAuthor, comments, nil
}

func (mr *MakananRepository) GetAllMakanan() ([]entity.MakananWithAuthor, error) {
	makanans := []entity.Makanan{}
	err := mr.db.Find(&makanans).Error
	if err != nil {
		return nil, err
	}

	makanansWithAuthor := []entity.MakananWithAuthor{}

	for _, makanan := range makanans {
		user := entity.User{}
		err := mr.db.Where("id = ?", makanan.UserID).First(&user).Error
		if err != nil {
			return nil, err
		}
		makanansWithAuthor = append(makanansWithAuthor, entity.MakananWithAuthor{
			Makanan:    makanan,
			Username:  user.Username,
			PhotoLink: user.ProfilePhotoLink,
		})
	}

	return makanansWithAuthor, nil
}

func (mr *MakananRepository) ContactMakanan(param model.MakananContact) (entity.Makanan, entity.User, error) {
	makanan := entity.Makanan{}
	err := mr.db.Where("id = ?", param.ID).First(&makanan).Error
	if err != nil {
		return entity.Makanan{}, entity.User{}, err
	}

	seller := entity.User{}
	err = mr.db.Where("id = ?", makanan.UserID).First(&seller).Error
	if err != nil {
		return entity.Makanan{}, entity.User{}, err
	}

	history := entity.History{
		ID:             uuid.New(),
		PostID:         makanan.ID,
		TypeOfPost:     "makanan",
		UserID:         param.AskerID,
		Title:          makanan.Title,
		SellerUsername: seller.Username,
		Price:          makanan.Price,
	}

	err = mr.db.Create(&history).Error
	if err != nil {
		return entity.Makanan{}, entity.User{}, err
	}

	return makanan, seller, nil
}
