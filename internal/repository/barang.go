package repository

import (
	"github.com/google/uuid"
	"github.com/kmdavidds/mager-spot-api/entity"
	"github.com/kmdavidds/mager-spot-api/model"
	"gorm.io/gorm"
)

type IBarangRepository interface {
	CreateBarang(barang entity.Barang) (entity.Barang, error)
	GetBarang(param model.BarangParam) (entity.BarangWithAuthor, []entity.Comment, error)
	GetAllBarang() ([]entity.BarangWithAuthor, error)
	ContactBarang(param model.BarangContact) (entity.Barang, entity.User, error)
}

type BarangRepository struct {
	db *gorm.DB
}

func NewBarangRepository(db *gorm.DB) IBarangRepository {
	return &BarangRepository{
		db: db,
	}
}

func (br *BarangRepository) CreateBarang(barang entity.Barang) (entity.Barang, error) {
	err := br.db.Create(&barang).Error
	if err != nil {
		return barang, err
	}

	return barang, nil
}

func (br *BarangRepository) GetBarang(param model.BarangParam) (entity.BarangWithAuthor, []entity.Comment, error) {

	barang := entity.Barang{}
	err := br.db.Where(&param).First(&barang).Error
	if err != nil {
		return entity.BarangWithAuthor{}, nil, err
	}

	user := entity.User{}
	err = br.db.Where("id = ?", barang.UserID).First(&user).Error
	if err != nil {
		return entity.BarangWithAuthor{}, nil, err
	}

	barangWithAuthor := entity.BarangWithAuthor{
		Barang:    barang,
		Username:  user.Username,
		PhotoLink: user.ProfilePhotoLink,
	}

	comments := []entity.Comment{}
	err = br.db.Where("post_id = ?", barang.ID).Find(&comments).Error
	if err != nil {
		return barangWithAuthor, nil, err
	}

	return barangWithAuthor, comments, nil
}

func (br *BarangRepository) GetAllBarang() ([]entity.BarangWithAuthor, error) {
	barangs := []entity.Barang{}
	err := br.db.Find(&barangs).Error
	if err != nil {
		return nil, err
	}

	barangsWithAuthor := []entity.BarangWithAuthor{}

	for _, barang := range barangs {
		user := entity.User{}
		err := br.db.Where("id = ?", barang.UserID).First(&user).Error
		if err != nil {
			return nil, err
		}
		barangsWithAuthor = append(barangsWithAuthor, entity.BarangWithAuthor{
			Barang:    barang,
			Username:  user.Username,
			PhotoLink: user.ProfilePhotoLink,
		})
	}

	return barangsWithAuthor, nil
}

func (br *BarangRepository) ContactBarang(param model.BarangContact) (entity.Barang, entity.User, error) {
	barang := entity.Barang{}
	err := br.db.Where("id = ?", param.ID).First(&barang).Error
	if err != nil {
		return entity.Barang{}, entity.User{}, err
	}

	user := entity.User{}
	err = br.db.Where("id = ?", barang.UserID).First(&user).Error
	if err != nil {
		return entity.Barang{}, entity.User{}, err
	}

	history := entity.History{
		ID: uuid.New(),
		PostID: barang.ID,
		UserID: param.AskerID,
	}

	err = br.db.Create(&history).Error
	if err != nil {
		return entity.Barang{}, entity.User{}, err
	}

	return barang, user, nil
}
