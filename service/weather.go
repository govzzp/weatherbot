package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"weather-bot/model"
)

func GetWeather(city string, lng, lat float64, token string) (*model.Weather, string, error) {

	url := fmt.Sprintf(
		"https://api.caiyunapp.com/v2.6/%s/%f,%f/weather",
		token, lng, lat,
	)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, "", err
	}

	q := req.URL.Query()
	q.Add("alert", "true")
	q.Add("dailysteps", "1")
	q.Add("hourlysteps", "24")
	req.URL.RawQuery = q.Encode()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()

	var res map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&res)

	result := res["result"].(map[string]interface{})
	daily := result["daily"].(map[string]interface{})

	temp := daily["temperature"].([]interface{})[0].(map[string]interface{})
	sky := daily["skycon"].([]interface{})[0].(map[string]interface{})

	w := &model.Weather{
		City:    city,
		Date:    time.Now().Format("2006-01-02"),
		MinTemp: temp["min"].(float64),
		MaxTemp: temp["max"].(float64),
		Sky:     sky["value"].(string),
	}

	msg := fmt.Sprintf("%s：%s %.1f~%.1f℃",
		city, w.Sky, w.MinTemp, w.MaxTemp)

	return w, msg, nil
}
