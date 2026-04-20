package router

import (
	"weather-bot/handler"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {

	r := gin.Default()

	r.GET("/history", handler.GetHistory(db))
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	return r
}
