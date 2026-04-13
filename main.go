package main

import (
    "github.com/gin-gonic/gin"
    "weather-bot/config"
	"weather-bot/dao"
	"weather-bot/model"
	"weather-bot/router"
	"weather-bot/service"

	"github.com/robfig/cron/v3"
)

func main() {
    gin.SetMode(gin.ReleaseMode)

	cfg := config.LoadConfig()

	db := dao.InitDB(cfg.Mysql.DSN)
	db.AutoMigrate(&model.Weather{})

	// 定时任务
	c := cron.New(cron.WithSeconds())

	c.AddFunc("0 30 07 * * *", func() {
		service.RunJob(cfg, db)
	})

	c.Start()

	r := router.SetupRouter(db)

	r.Run(":14250")
}
