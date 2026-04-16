package service

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"
	"weather-bot/util"

	"github.com/goccy/go-json"
)

func SendFeishu(webhook string, msg FeishuMessage) error {
	util.Log.Info("feishu request sent")
	jsonData, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", webhook, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println("飞书返回:", string(body))

	// ❗ HTTP 层错误
	if resp.StatusCode != 200 {
		return fmt.Errorf("feishu http error: %d, body=%s", resp.StatusCode, string(body))
	}

	// ❗ 业务层错误
	var res struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	}

	if err := json.Unmarshal(body, &res); err != nil {
		return fmt.Errorf("invalid feishu response: %s", string(body))
	}

	if res.Code != 0 {
		return fmt.Errorf("feishu biz error: code=%d msg=%s", res.Code, res.Msg)
	}

	return nil
}

var feishuLimiter = time.NewTicker(1 * time.Second)
var feishuCh = make(chan struct{}, 1)

func init() {
	go func() {
		for range feishuLimiter.C {
			select {
			case feishuCh <- struct{}{}:
			default:
			}
		}
	}()
}

// SafeSendFeishu 增加飞书的限流代码
func SafeSendFeishu(webhook string, msg FeishuMessage) error {

	<-feishuCh // 🚀 限流点

	return SendFeishu(webhook, msg)
}
