package postgresql

import (
	"log"

	"github.com/kmdavidds/mager-spot-api/entity"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&entity.User{},
		&entity.Barang{},
		&entity.Kos{},
		&entity.Makanan{},
		&entity.Ojek{},
		&entity.Comment{},
	)

	if err != nil {
		log.Fatalf("failed migration db: %v", err)
	}
}