package service

import (
	"strings"
	"weather-bot/model"
)

type WeatherMood struct {
	Level   string // SUNNY / NORMAL / RAIN / ALERT
	Color   string // green / blue / red / orange
	Message string
}

func GetWeatherMood(w model.SimpleWeather) WeatherMood {

	// ⛔ 高风险：空气污染
	if w.AQI >= 150 {
		return WeatherMood{
			Level:   "ALERT",
			Color:   "red",
			Message: "空气污染较严重，建议减少外出",
		}
	}

	// 🌧 强降雨
	if strings.Contains(w.RainHint, "90") || strings.Contains(w.RainHint, "85") {
		return WeatherMood{
			Level:   "ALERT",
			Color:   "red",
			Message: "强降雨概率较高，请谨慎出行",
		}
	}

	// 🌧 可能下雨
	if strings.Contains(w.RainHint, "50") {
		return WeatherMood{
			Level:   "RAIN",
			Color:   "blue",
			Message: "可能降雨，建议带伞",
		}
	}

	// ☀️ 晴天
	if w.Sky == "晴" || w.Sky == "少云" {
		return WeatherMood{
			Level:   "SUNNY",
			Color:   "green",
			Message: "天气晴朗，适合出行",
		}
	}

	// 🌤 默认
	return WeatherMood{
		Level:   "NORMAL",
		Color:   "blue",
		Message: "天气较为稳定，可正常出行",
	}
}
