package models

import (
	"tugas-rest-api/models/games"
	"tugas-rest-api/models/kepemilikan"
	"tugas-rest-api/models/users"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/steam"))
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&users.Users{}, &games.Games{}, &kepemilikan.Kepemilikan{})

	DB = database

	// Tampilkan isi DB

}
