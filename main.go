//go:generate swag init
//go:generate go run ./cmd/migrate
//go:generate go run ./cmd/gen
package main

import (
	"log"
	"pastebingo/dal/models"
	"pastebingo/dal/query"
	"pastebingo/internal/server"

	swaggerFiles "github.com/swaggo/files"

	_ "pastebingo/docs" // 这里导入生成的 docs 包

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title			Pastebingo API
// @version		1.0
// @description	This is a sample server for Pastebingo.
// @host			localhost:8080
// @BasePath		/
func main() {
	models.Migrate()
	db := models.DB
	query.SetDefault(db)

	r := gin.Default()
	server.RegisterRoutes(r)
	// Swagger 路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 启动服务器
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
