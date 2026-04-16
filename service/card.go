package service

import (
	"fmt"
	"strings"
	"weather-bot/model"
)

//func BuildFeishuCard(w model.SimpleWeather) map[string]interface{} {
//
//	mood := GetWeatherMood(w)
//
//	header := fmt.Sprintf("🌤 %s · %s", w.City, w.Date)
//
//	content := fmt.Sprintf(
//		"🌡 温度：%.1f℃ ~ %.1f℃\n☁ 天气：%s\n💧 湿度：%d%%\n🌬 风速：%.1f km/h\n\n💡 %s",
//		w.MinTemp, w.MaxTemp, w.Sky, w.Humidity, w.WindSpeed,
//		mood.Message,
//	)
//
//	return map[string]interface{}{
//		"msg_type": "interactive",
//		"card": map[string]interface{}{
//			"header": map[string]interface{}{
//				"title": map[string]string{
//					"tag":     "plain_text",
//					"content": header,
//				},
//				"template": mood.Color, // ⭐ 核心：动态颜色
//			},
//
//			"elements": []interface{}{
//
//				// 主体
//				map[string]interface{}{
//					"tag": "div",
//					"text": map[string]string{
//						"tag":     "lark_md",
//						"content": content,
//					},
//				},
//
//				map[string]interface{}{"tag": "hr"},
//
//				// 额外信息
//				map[string]interface{}{
//					"tag": "div",
//					"text": map[string]string{
//						"tag":     "lark_md",
//						"content": fmt.Sprintf("🌈 情绪等级：%s", mood.Level),
//					},
//				},
//			},
//		},
//	}
//}

func BuildFeishuCard(w model.SimpleWeather) FeishuMessage {

	mood := GetWeatherMood(w)

	header := fmt.Sprintf("🌤 %s · %s", w.City, w.Date)

	body := fmt.Sprintf(
		"🌡 温度：%.1f℃ ~ %.1f℃\n☁ 天气：%s\n💧 湿度：%d%%\n🌬 风速：%.1f km/h\n🌧 降雨概率：%.0f%%\n\n💡 %s",
		w.MinTemp,
		w.MaxTemp,
		w.Sky,
		w.Humidity,
		w.WindSpeed,
		w.RainProb*100,
		mood.Message,
	)

	return FeishuMessage{
		MsgType: "interactive",
		Card: FeishuCard{
			Config: CardConfig{
				WideScreenMode: true,
			},
			Header: CardHeader{
				Title: TextModule{
					Tag:     "plain_text",
					Content: header,
				},
				Template: mood.Color,
			},
			Elements: []Element{
				{
					Tag: "div",
					Text: &TextModule{
						Tag:     "lark_md",
						Content: body,
					},
				},
				{Tag: "hr"},
				{
					Tag: "div",
					Text: &TextModule{
						Tag:     "lark_md",
						Content: fmt.Sprintf("🌈 风险等级：%s（score=%d）", mood.Level, mood.Score),
					},
				},
				{
					Tag: "div",
					Text: &TextModule{
						Tag:     "lark_md",
						Content: fmt.Sprintf("🏷 标签：%s", strings.Join(mood.Tags, ", ")),
					},
				},
			},
		},
	}
}
