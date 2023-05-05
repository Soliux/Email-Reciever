package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func ConnectDB() error {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=myuser dbname=mydb password=mypassword sslmode=disable")
	if err != nil {
		return err
	}

	DB = db
	return nil
}
