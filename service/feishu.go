package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func SendFeishu(webhook string, msgs []string) {

	content := ""
	for _, m := range msgs {
		content += fmt.Sprintf("• %s\n", m)
	}

	body := map[string]interface{}{
		"msg_type": "interactive",
		"card": map[string]interface{}{
			"header": map[string]interface{}{
				"title": map[string]string{
					"tag":     "plain_text",
					"content": "🌤 每日天气",
				},
			},
			"elements": []map[string]interface{}{
				{
					"tag": "div",
					"text": map[string]string{
						"tag":     "lark_md",
						"content": fmt.Sprintf("📅 %s\n\n%s", time.Now().Format("2006-01-02"), content),
					},
				},
			},
		},
	}

	data, _ := json.Marshal(body)
	fmt.Println(webhook)
	fmt.Println(string(data))
	http.Post(webhook, "application/json", bytes.NewBuffer(data))

}
