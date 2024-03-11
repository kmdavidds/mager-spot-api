package postgresql

import (
	"log"

	"github.com/kmdavidds/mager-spot-api/entity"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&entity.User{},
		&entity.ApartmentPost{},
		&entity.FoodPost{},
		&entity.ProductPost{},
		&entity.ShuttlePost{},
		&entity.Comment{},
		&entity.History{},
	)

	if err != nil {
		log.Fatalf("failed migration db: %v", err)
	}
}