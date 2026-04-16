package service

import (
	"fmt"
	"weather-bot/config"
	"weather-bot/model"

	"gorm.io/gorm"
)

func RunJob(cfg *config.Config, db *gorm.DB) {

	for _, c := range cfg.Cities {

		data, err := GetWeatherRaw(c.Name, c.Lng, c.Lat, cfg.Caiyun.Token)
		if err != nil {
			fmt.Println("API错误:", err)
			continue
		}

		w := ParseWeather(c.Name, data)

		err = db.Where("city = ? AND date = ?", w.City, w.Date).
			Assign(model.SimpleWeather{
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
			}).
			FirstOrCreate(&model.SimpleWeather{
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
			}).Error

		if err != nil {
			fmt.Println("数据库错误:", err)
		}
		// 飞书
		// ✅ 飞书
		card := BuildFeishuCard(w)
		if err := SafeSendFeishu(cfg.Feishu.Webhook, card); err != nil {
			fmt.Println("飞书发送失败:", err)
		}
	}
}
