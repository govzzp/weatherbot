package service

import (
	"fmt"
	"strings"
	"weather-bot/model"
)

func GetWeatherLevel(w model.SimpleWeather) string {

	// 🔴 1. 官方预警（最高优先级）
	if w.Alert != "" {
		return "red"
	}

	// 🔴 2. AQI
	if w.AQI >= 150 {
		return "red"
	}
	if w.AQI >= 100 {
		return "orange"
	}

	// 🔴 3. 极端天气关键词
	if strings.Contains(w.Sky, "雷") ||
		strings.Contains(w.Sky, "沙") ||
		strings.Contains(w.Sky, "暴") {
		return "red"
	}

	// 🌧 4. 降雨概率解析
	if strings.Contains(w.RainHint, "概率") {
		var p int
		_, err := fmt.Sscanf(w.RainHint, "%*s概率%d%%", &p)
		if err != nil {
			return ""
		}

		if p >= 85 {
			return "red"
		}
		if p >= 60 {
			return "orange"
		}
		if p >= 30 {
			return "blue"
		}
	}

	// ☀️ 5. 默认晴天
	if w.Sky == "晴" {
		return "green"
	}

	return "blue"
}
