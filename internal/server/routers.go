package server

import (
	"github.com/gin-gonic/gin"
	"pastebingo/internal/api"
)

func RegisterRoutes(router *gin.Engine) {

	ApiV1 := router.Group("/v1")
	api.UserCreate(ApiV1)

}
