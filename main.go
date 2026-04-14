package main

import (
	"weather-bot/config"
	"weather-bot/dao"
	"weather-bot/router"
	"weather-bot/service"

	"github.com/robfig/cron/v3"
)

func main() {
	//gin.SetMode(gin.ReleaseMode)

	cfg := config.LoadConfig()
	db := dao.Getdb()
	// 定时任务
	c := cron.New(cron.WithSeconds())

	_, err := c.AddFunc("00 30 07 * * *", func() {

		service.RunJob(cfg, db)
	})
	if err != nil {
		return
	}

	c.Start()
	go service.RunJob(cfg, db)

	r := router.SetupRouter(db)

	panic(r.Run(":14250"))
}
