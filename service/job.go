package service

import (
	"go.uber.org/zap"
	"weather-bot/config"
	"weather-bot/model"
	"weather-bot/util"

	"gorm.io/gorm"
)

func RunJob(cfg *config.Config, db *gorm.DB) {

	util.Log.Info("job start")
	for _, c := range cfg.Cities {
		util.Log.Info("weather job start")

		util.Log.Info("fetch weather",
			zap.String("city", c.Name),
		)
		data, err := GetWeatherRaw(c.Name, c.Lng, c.Lat, cfg.Caiyun.Token)
		if err != nil {
			util.Log.Error("weather api failed",
				zap.String("city", c.Name),
				zap.Error(err),
			)
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
				RainProb:    w.RainProb,
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
				RainProb:    w.RainProb,
			}).Error

		if err != nil {
			util.Log.Error("weather insert data failed", zap.Error(err))
		}
		util.Log.Info("db upsert success",
			zap.String("city", w.City),
			zap.String("date", w.Date),
		)
		// 飞书
		// ✅ 飞书
		card := BuildFeishuCard(w)
		util.Log.Info("send feishu",
			zap.String("city", w.City),
		)
		if err := SafeSendFeishu(cfg.Feishu.Webhook, card); err != nil {
			util.Log.Warn("feishu failed",
				zap.Error(err),
			)
		}
	}
}
