package service

import (
	"fmt"
	"weather-bot/config"
	"weather-bot/model"

	"gorm.io/gorm"
)

func RunJob(cfg *config.Config, db *gorm.DB) {

	for _, c := range cfg.Cities {

		raw, err := GetWeatherRaw(c.Name, c.Lng, c.Lat, cfg.Caiyun.Token)
		if err != nil {
			fmt.Println("API错误:", err)
			continue
		}

		w, err := ParseWeather(c.Name, raw)
		if err != nil {
			fmt.Println("解析错误:", err)
			continue
		}

		db.Where(model.SimpleWeather{
			City: w.City,
			Date: w.Date,
		}).FirstOrCreate(&model.SimpleWeather{
			City:        w.City,
			Date:        w.Date,
			MinTemp:     w.MinTemp,
			MaxTemp:     w.MaxTemp,
			Sky:         w.Sky,
			Humidity:    w.Humidity,
			WindSpeed:   w.WindSpeed,
			AQI:         w.AQI,
			AQIDesc:     w.AQIDesc,
			FeelingTemp: w.FeelingTemp,
			Alert:       w.Alert,
			RainHint:    w.RainHint,
		})
		// 飞书
		card := BuildFeishuCard(SimpleWeather(w))
		SendFeishu(cfg.Feishu.Webhook, card)
	}
}
