package util

func FormatSky(s string) string {
	m := map[string]string{
		"CLEAR_DAY":           "晴",
		"CLOUDY":              "阴",
		"PARTLY_CLOUDY_DAY":   "多云",
		"PARTLY_CLOUDY_NIGHT": "多云",
		"LIGHT_RAIN":          "小雨",
		"MODERATE_RAIN":       "中雨",
	}
	if v, ok := m[s]; ok {
		return v
	}
	return s
}
