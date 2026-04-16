package handler

import (
	"fmt"
	"weather-bot/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetHistory(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		city := c.Query("city")

		var list []model.SimpleWeather
		db.Where("city = ?", city).
			Order("date desc").
			Limit(7).
			Find(&list)

		// ⭐ 不再直接返回 JSON，而是“结构化卡片数据”
		var result []map[string]interface{}

		for _, w := range list {
			result = append(result, map[string]interface{}{
				"date":    w.Date,
				"weather": w.Sky,
				"temp":    fmt.Sprintf("%.1f~%.1f℃", w.MinTemp, w.MaxTemp),
				"aqi":     w.AQI,
			})
		}

		c.JSON(200, gin.H{
			"city": city,
			"list": result,
		})
	}
}
