package main

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("db.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	g := gen.NewGenerator(gen.Config{
		OutPath: "dal/query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	g.UseDB(db)

	g.ApplyBasic(g.GenerateModel("users"), g.GenerateModel("posts"))

	g.Execute()

}
