package main

import (
	"pastebingo/dal/models"

	"gorm.io/gen"
)

func main() {
	db := models.DB

	g := gen.NewGenerator(gen.Config{
		OutPath: "dal/query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	g.UseDB(db)

	g.ApplyBasic(g.GenerateModel("users"), g.GenerateModel("posts"))

	g.Execute()

}
