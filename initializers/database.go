package initializers

import (
	"log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() (*gorm.DB, error) {
	if DB != nil {
		return DB, nil
	}

	var err error
	DB, err = gorm.Open(sqlite.Open("auth.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Error Connect to Database: ", err)
		return nil, err
	}
	return DB, nil
}

