package service

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"
	"weather-bot/model"

	"github.com/goccy/go-json"
)

func BuildFeishuCard(w model.SimpleWeather) map[string]interface{} {

	header := fmt.Sprintf("🌤 %s · %s", w.City, w.Date)

	mainInfo := fmt.Sprintf(
		"🌡 **温度**：`%.1f℃ ~ %.1f℃`\n"+
			"☁ **天气**：**%s**\n"+
			"💧 **湿度**：%d%%\n"+
			"🌬 **风速**：%.1f km/h",
		w.MinTemp, w.MaxTemp, w.Sky, w.Humidity, w.WindSpeed,
	)

	extraInfo := fmt.Sprintf(
		"🌈 **体感温度**：%.1f℃\n"+
			"🌿 **空气质量**：%s (AQI %d)",
		w.FeelingTemp, w.AQIDesc, w.AQI,
	)

	return map[string]interface{}{
		"msg_type": "interactive",
		"card": map[string]interface{}{
			"header": map[string]interface{}{
				"title": map[string]string{
					"tag":     "plain_text",
					"content": header,
				},
				"template": "blue",
			},
			"elements": []interface{}{

				// 🌟 核心天气
				map[string]interface{}{
					"tag": "div",
					"text": map[string]string{
						"tag":     "lark_md",
						"content": mainInfo,
					},
				},

				map[string]interface{}{"tag": "hr"},

				// 🌿 辅助信息
				map[string]interface{}{
					"tag": "div",
					"text": map[string]string{
						"tag":     "lark_md",
						"content": extraInfo,
					},
				},

				map[string]interface{}{"tag": "hr"},

				// 📢 天气提示
				map[string]interface{}{
					"tag": "div",
					"text": map[string]string{
						"tag":     "lark_md",
						"content": "📢 **天气提示**\n" + w.Alert,
					},
				},

				// 🌧 降雨提醒
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
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	body, _ := io.ReadAll(resp.Body)

	fmt.Println("飞书返回:", string(body))

	// ✅ 检查 HTTP 状态码
	if resp.StatusCode != 200 {
		return fmt.Errorf("飞书请求失败: status=%d", resp.StatusCode)
	}

	// ✅ 解析返回 JSON
	var res map[string]interface{}
	_ = json.Unmarshal(body, &res)

	if code, ok := res["code"].(float64); ok && code != 0 {
		return fmt.Errorf("飞书错误: %v", res)
	}

	return nil
}

// 增加飞书的限流代码
func SafeSendFeishu(webhook string, card map[string]interface{}) error {
	err := SendFeishu(webhook, card)
	time.Sleep(1 * time.Second)
	return err
}
