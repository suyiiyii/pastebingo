package main

import (
	"log"

	"pastebingo/dal/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("db.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal(err)
	}
	err = db.AutoMigrate(&models.Post{})
	if err != nil {
		log.Fatal(err)
	}

}