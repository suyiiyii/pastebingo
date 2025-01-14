//go:generate swag init --parseDependency --parseInternal
//go:generate swag fmt
//go:generate go run ./cmd/migrate
//go:generate go run ./cmd/gen
package main

import (
	"log"
	"net/http"
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
// @BasePath		/v1/
func main() {
	models.Migrate()
	db := models.DB
	query.SetDefault(db)

	r := gin.Default()
	server.RegisterRoutes(r)
	// Swagger 路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 挂载前端页面
	// 定义文件服务器
	fileServer := http.FileServer(http.Dir("./dist"))

	// 去除 URL 前缀
	r.GET("/", func(c *gin.Context) {
		http.ServeFile(c.Writer, c.Request, "./dist/index.html")
	})
	// 处理其他静态文件
	r.NoRoute(func(c *gin.Context) {
		fileServer.ServeHTTP(c.Writer, c.Request)
	})

	// 启动服务器
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
