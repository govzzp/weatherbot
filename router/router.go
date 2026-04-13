package router

import (
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "weather-bot/handler"
)

func SetupRouter(db *gorm.DB) *gin.Engine {

    r := gin.Default()

    r.GET("/history", handler.GetHistory(db))

    return r
}