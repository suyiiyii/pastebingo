package models

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

const DB_ENV = "GORM_DB"

func initDB() {
	err := godotenv.Load()
	if err != nil {
		log.Default().Println("Error loading .env file")
	}
	dbName := os.Getenv(DB_ENV)
	if dbName == "" {
		dbName = "default.db"
	}
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = db
}

func init() {
	initDB()
}
