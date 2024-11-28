package main

import (
	"log"
	"pastebingo/dal/query"
	"pastebingo/internal/server"

	swaggerFiles "github.com/swaggo/files"

	_ "pastebingo/docs" // 这里导入生成的 docs 包

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var GetQ func() *query.Query

// @title			Pastebingo API
// @version		1.0
// @description	This is a sample server for Pastebingo.
// @host			localhost:8080
// @BasePath		/
func main() {
	db, err := gorm.Open(sqlite.Open("db.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	Q := query.Use(db)
	query.SetDefault(db)

	GetQ = func() *query.Query {
		return Q
	}

	r := gin.Default()
	server.RegisterRoutes(r)
	// Swagger 路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 启动服务器
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
