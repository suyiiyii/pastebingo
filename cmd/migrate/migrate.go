package main

import (
	"log"
	"pastebingo/dal/models"
)

// use pre-defined models to create tables
func main() {
	db := models.DB

	err := db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal(err)
	}
	err = db.AutoMigrate(&models.Post{})
	if err != nil {
		log.Fatal(err)
	}

}
