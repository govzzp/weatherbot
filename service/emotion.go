package service

import (
	"strings"
	"weather-bot/model"
)

type WeatherMood struct {
	Level   string
	Color   string
	Message string

	Score int
	Tags  []string
}

func GetWeatherMood(w model.SimpleWeather) WeatherMood {

	score := 0
	tags := make([]string, 0)

	// 🌫 AQI 分级
	switch {
	case w.AQI >= 200:
		score -= 4
		tags = append(tags, "severe_pollution")
	case w.AQI >= 150:
		score -= 3
		tags = append(tags, "pollution")
	case w.AQI >= 100:
		score -= 2
		tags = append(tags, "light_pollution")
	}

	// 🌧 降雨概率（结构化）
	switch {
	case w.RainProb >= 0.85:
		score -= 4
		tags = append(tags, "heavy_rain")
	case w.RainProb >= 0.5:
		score -= 2
		tags = append(tags, "rain")
	case w.RainProb >= 0.2:
		score -= 1
	}

	// ☀️ 天气状态
	if strings.Contains(w.Sky, "晴") {
		score += 2
		tags = append(tags, "sunny")
	}
	if strings.Contains(w.Sky, "云") {
		score += 0
	}
	if strings.Contains(w.Sky, "雨") {
		score -= 1
	}

	// 🌬 风速
	if w.WindSpeed > 30 {
		score -= 2
		tags = append(tags, "windy")
	}

	var level, color, message string

	switch {
	case score <= -5:
		level = "ALERT"
		color = "red"
		message = "极端天气风险较高，建议减少外出"

	case score <= -2:
		level = "WARN"
		color = "orange"
		message = "天气较差，出行请注意安全"

	case score <= 1:
		level = "NORMAL"
		color = "blue"
		message = "天气一般，可正常出行"

	default:
		level = "GOOD"
		color = "green"
		message = "天气良好，适合出行"
	}

	return WeatherMood{
		Level:   level,
		Color:   color,
		Message: message,
		Score:   score,
		Tags:    tags,
	}
}
