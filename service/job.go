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

		// DB（你原逻辑不动）
		db.Where("city = ? AND date = ?", w.City, w.Date).
			Assign(model.SimpleWeather{
				MinTemp: w.MinTemp,
				MaxTemp: w.MaxTemp,
			}).
			FirstOrCreate(&model.SimpleWeather{})

		// 飞书
		card := BuildFeishuCard(SimpleWeather(w))
		SendFeishu(cfg.Feishu.Webhook, card)
	}
}
