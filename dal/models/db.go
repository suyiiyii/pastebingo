package models

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

const DB_ENV = "GORM_DB"

func initDB() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(sqlite.Open(os.Getenv(DB_ENV)), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = db
}

func init() {
	initDB()
}
