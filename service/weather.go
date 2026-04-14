package service

import (
	"fmt"
	"io"
	"net/http"
	"time"
	"weather-bot/model"
	"weather-bot/util"

	"github.com/goccy/go-json"
)

func GetWeatherRaw(city string, lng, lat float64, token string) (map[string]interface{}, error) {

	url := fmt.Sprintf(
		"https://api.caiyunapp.com/v2.6/%s/%f,%f/weather",
		token, lng, lat,
	)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("alert", "true")
	q.Add("dailysteps", "1")
	q.Add("hourlysteps", "24")
	req.URL.RawQuery = q.Encode()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)

	var res map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}

	// ⭐ 防止 token invalid
	if res["status"] != "ok" {
		return nil, fmt.Errorf("api error: %v", res["status"])
	}

	return res, nil
}
func ParseWeather(city string, data map[string]interface{}) (model.SimpleWeather, error) {

	result, ok := data["result"].(map[string]interface{})
	if !ok {
		return model.SimpleWeather{}, fmt.Errorf("invalid result")
	}

	realtime := result["realtime"].(map[string]interface{})
	daily := result["daily"].(map[string]interface{})

	temp0 := daily["temperature"].([]interface{})[0].(map[string]interface{})

	aq := realtime["air_quality"].(map[string]interface{})
	aqi := aq["aqi"].(map[string]interface{})

	// 🌧 降雨逻辑
	rainHint := "无明显降雨"

	if hourly, ok := result["hourly"].(map[string]interface{}); ok {
		if precList, ok := hourly["precipitation"].([]interface{}); ok {
			for _, v := range precList {
				item := v.(map[string]interface{})

				if p, ok := item["probability"].(float64); ok && p >= 50 {
					rainHint = fmt.Sprintf("%s 有降雨（%.0f%%）",
						item["datetime"].(string)[11:16],
						p,
					)
					break
				}
			}
		}
	}

	return model.SimpleWeather{
		City:        city,
		Date:        time.Now().Format("2006-01-02"),
		MinTemp:     temp0["min"].(float64),
		MaxTemp:     temp0["max"].(float64),
		Sky:         util.FormatSky(realtime["skycon"].(string)),
		Humidity:    int(realtime["humidity"].(float64) * 100),
		WindSpeed:   realtime["wind"].(map[string]interface{})["speed"].(float64),
		AQI:         int(aqi["chn"].(float64)),
		AQIDesc:     aq["description"].(map[string]interface{})["chn"].(string),
		FeelingTemp: realtime["apparent_temperature"].(float64),
		Alert:       result["forecast_keypoint"].(string),
		RainHint:    rainHint,
	}, nil
}
