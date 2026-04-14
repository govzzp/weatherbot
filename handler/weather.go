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

		query := db.Order("date asc")

		// ✅ 有城市才加条件
		if city != "" {
			query = query.Where("city = ?", city)
		}

		query.Find(&list)

		c.JSON(http.StatusOK, list)
	}
}
