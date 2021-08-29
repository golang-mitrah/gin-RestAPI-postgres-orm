package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	connection, err := gorm.Open(postgres.Open("host=localhost user=sm dbname=gin_goorm_rest port=5432 password="), &gorm.Config{})
	if err != nil {
		panic("could not connect to DB!")
	}

	DB = connection

}
