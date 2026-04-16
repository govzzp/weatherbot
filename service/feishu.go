package service

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/goccy/go-json"
)

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

// SafeSendFeishu 增加飞书的限流代码
func SafeSendFeishu(webhook string, card map[string]interface{}) error {
	err := SendFeishu(webhook, card)
	time.Sleep(1 * time.Second)
	return err
}
