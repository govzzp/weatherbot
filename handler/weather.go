package handler

import (
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "weather-bot/model"
)

func GetHistory(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {

        city := c.Query("city")

        var list []model.Weather
        db.Where("city = ?", city).Order("date asc").Find(&list)

        c.JSON(200, list)
    }
}