package service

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/goccy/go-json"
)

func BuildFeishuCard(w SimpleWeather) map[string]interface{} {

	content := fmt.Sprintf(
		"📅 %s\n\n🌡 **温度**：%.1f℃ ~ %.1f℃\n☁ **天气**：%s\n💧 **湿度**：%d%%\n🌬 **风速**：%.1f km/h",
		w.Date, w.MinTemp, w.MaxTemp, w.Sky, w.Humidity, w.WindSpeed,
	)

	return map[string]interface{}{
		"msg_type": "interactive",
		"card": map[string]interface{}{
			"header": map[string]interface{}{
				"title": map[string]string{
					"tag":     "plain_text",
					"content": "🌤 每日天气 · " + w.City,
				},
				"template": "blue",
			},
			"elements": []interface{}{
				map[string]interface{}{
					"tag": "div",
					"text": map[string]string{
						"tag":     "lark_md",
						"content": content,
					},
				},
				map[string]interface{}{
					"tag": "hr",
				},
				map[string]interface{}{
					"tag": "div",
					"fields": []interface{}{
						map[string]interface{}{
							"is_short": true,
							"text": map[string]string{
								"tag":     "lark_md",
								"content": fmt.Sprintf("🌈 **体感温度**\n%.1f℃", w.FeelingTemp),
							},
						},
						map[string]interface{}{
							"is_short": true,
							"text": map[string]string{
								"tag":     "lark_md",
								"content": fmt.Sprintf("🌿 **空气质量**\n%s (AQI %d)", w.AQIDesc, w.AQI),
							},
						},
					},
				},
				map[string]interface{}{
					"tag": "hr",
				},
				map[string]interface{}{
					"tag": "div",
					"text": map[string]string{
						"tag":     "lark_md",
						"content": "📝 **天气提示**\n" + w.Alert,
					},
				},
				map[string]interface{}{
					"tag": "div",
					"text": map[string]string{
						"tag":     "lark_md",
						"content": "🌧 **降雨提醒**\n" + w.RainHint,
					},
				},
			},
		},
	}
}

func SendFeishu(webhook string, card map[string]interface{}) error {

	jsonData, err := json.Marshal(card)
	if err != nil {
		return err
	}

	resp, err := http.Post(webhook, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println("飞书返回:", string(body))

	return nil
}
