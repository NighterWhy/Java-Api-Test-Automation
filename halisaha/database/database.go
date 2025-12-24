package database

import (
	"halisaha/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	dsn := "host=localhost user=serhatsimsek password=123456 dbname=halisaha_db port=5433 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed while connected to database!", err)

	}

	database.AutoMigrate(&models.Rezervasyon{})
	database.AutoMigrate(&models.Saha{})
	database.AutoMigrate(&models.User{})

	DB = database
	log.Println("Veritabanına Bağlantı Başarılı")

}
