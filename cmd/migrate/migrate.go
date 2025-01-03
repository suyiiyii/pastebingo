package main

import (
	"pastebingo/dal/models"
)

// use pre-defined models to create tables
func main() {
	models.Migrate()

}
