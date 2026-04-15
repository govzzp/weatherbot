package util

func FormatSky(s string) string {
	m := map[string]string{
		// 晴天
		"CLEAR_DAY":   "☀️ 晴",
		"CLEAR_NIGHT": "🌙 晴",

		// 多云
		"PARTLY_CLOUDY_DAY":   "⛅ 多云",
		"PARTLY_CLOUDY_NIGHT": "☁️ 多云",
		"CLOUDY":              "☁️ 阴",

		// 雨
		"LIGHT_RAIN":    "🌧 小雨",
		"MODERATE_RAIN": "🌧 中雨",
		"HEAVY_RAIN":    "🌧 大雨",
		"STORM_RAIN":    "⛈ 暴雨",

		// 雪
		"LIGHT_SNOW":    "❄️ 小雪",
		"MODERATE_SNOW": "❄️ 中雪",
		"HEAVY_SNOW":    "❄️ 大雪",
		"STORM_SNOW":    "❄️ 暴雪",

		// 特殊天气
		"FOG":  "🌫 雾",
		"HAZE": "🌫 霾",

		// 沙尘类（你刚问的重点）
		"DUST": "🌪 浮尘",
		"SAND": "🌪 沙尘",

		// 风
		"WIND": "💨 大风",
	}

	if v, ok := m[s]; ok {
		return v
	}
	return "未知天气(" + s + ")"
}
