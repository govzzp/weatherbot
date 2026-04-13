package service

import (
	"fmt"
	"gorm.io/gorm"
	"weather-bot/config"
)

func RunJob(cfg *config.Config, db *gorm.DB) {

	var msgs []string

	for _, c := range cfg.Cities {

		w, msg, err := GetWeather(c.Name, c.Lng, c.Lat, cfg.Caiyun.Token)
		if err != nil {
			fmt.Println("error:", err)
			continue
		}

		db.Where("city=? AND date=?", w.City, w.Date).
			FirstOrCreate(w)

		msgs = append(msgs, msg)
		fmt.Println(msgs)
	}

	SendFeishu(cfg.Feishu.Webhook, msgs)
}
