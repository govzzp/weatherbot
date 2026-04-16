package model

type SimpleWeather struct {
	City string
	Date string

	MinTemp float64
	MaxTemp float64

	Sky string

	Humidity    int
	WindSpeed   float64
	AQI         int
	AQIDesc     string
	FeelingTemp float64

	Alert    string
	RainHint string
}
type CaiyunResponse struct {
	Result struct {
		Realtime struct {
			Skycon      string  `json:"skycon"`
			Humidity    float64 `json:"humidity"`
			Temperature float64 `json:"temperature"`
			Wind        struct {
				Speed float64 `json:"speed"`
			} `json:"wind"`
			AirQuality struct {
				AQI struct {
					CHN int `json:"chn"`
				} `json:"aqi"`
				Description struct {
					CHN string `json:"chn"`
				} `json:"description"`
			} `json:"air_quality"`
		} `json:"realtime"`

		Daily struct {
			Temperature []struct {
				Min float64 `json:"min"`
				Max float64 `json:"max"`
			} `json:"temperature"`
		} `json:"daily"`

		ForecastKeypoint string `json:"forecast_keypoint"`
	} `json:"result"`
}
