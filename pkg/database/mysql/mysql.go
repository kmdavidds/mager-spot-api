package mysql

import (
	"log"

	"github.com/kmdavidds/mager-spot-api/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDatabase() *gorm.DB {
	db, err := gorm.Open(mysql.Open(config.MySQLConstructDSN()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		log.Fatalf("Error while connecting to database: %v", err)
		return nil
	}

	return db
}
