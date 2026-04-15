package router

import (
	"weather-bot/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {

	r := gin.Default()

	r.GET("/history", handler.GetHistory(db))

	return r
}
