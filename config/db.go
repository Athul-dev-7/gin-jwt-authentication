package config

import (
	"log"

	"jwt-authentication/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

/*	Here, we are defining an instance of the database.
	This variable will be used across the entire application to communicate with the database.
*/
var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=jwt password=jwt dbname=jwt port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Failed to connect DB %v\n", err)
		panic(err)
	}
	log.Println("Connected to Database!")
	db.AutoMigrate(&models.User{})
	DB = db
}
