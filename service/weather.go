package service

import (
	"fmt"
	"net/http"
	"time"
	"weather-bot/model"
	"weather-bot/util"

	"github.com/goccy/go-json"
)

func GetWeatherRaw(city string, lng, lat float64, token string) (*model.CaiyunResponse, error) {
	url := fmt.Sprintf("https://api.caiyunapp.com/v2.6/%s/%f,%f/weather", token, lng, lat)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var data model.CaiyunResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return &data, nil
}

func ParseWeather(city string, data *model.CaiyunResponse) model.SimpleWeather {

	rt := data.Result.Realtime
	daily := data.Result.Daily.Temperature[0]

	return model.SimpleWeather{
		City:        city,
		Date:        time.Now().Format("2006-01-02"),
		MinTemp:     daily.Min,
		MaxTemp:     daily.Max,
		Sky:         util.FormatSky(rt.Skycon),
		Humidity:    int(rt.Humidity * 100),
		WindSpeed:   rt.Wind.Speed,
		AQI:         rt.AirQuality.AQI.CHN,
		AQIDesc:     rt.AirQuality.Description.CHN,
		FeelingTemp: rt.Temperature,
		Alert:       data.Result.ForecastKeypoint,
	}
}
