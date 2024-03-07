package postgresql

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  os.Getenv("POSTGRESQL_DSN"),
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		log.Fatalf("Error while connecting to database: %v", err)
		return nil
	}

	return db
}
