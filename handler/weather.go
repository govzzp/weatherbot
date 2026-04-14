package handler

import (
	"net/http"
	"weather-bot/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetHistory(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		city := c.Query("city")

		var list []model.SimpleWeather
		db.Where("city = ?", city).Order("date asc").Find(&list)

		c.JSON(http.StatusOK, list)
	}
}
