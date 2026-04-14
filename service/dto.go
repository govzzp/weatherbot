package service

type SimpleWeather struct {
	City        string
	Date        string
	MinTemp     float64
	MaxTemp     float64
	Sky         string
	Humidity    int
	WindSpeed   float64
	AQI         int
	AQIDesc     string
	FeelingTemp float64
	Alert       string
	RainHint    string
}
