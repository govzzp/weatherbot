package dao

import (
	"fmt"
	"os"
	"time"
	"weather-bot/model"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	if err := initDB(); err != nil {
		panic(err)
	}
}
func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
func initDB() (err error) {
	InitConfig()
	host := viper.GetString("database.host")
	port := viper.GetString("database.port")
	username := viper.GetString("database.username")
	password := viper.GetString("database.password")
	database := viper.GetString("database.database")
	charset := viper.GetString("database.charset")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local", username, password, host, port, database, charset)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	// ⭐ 关键配置
	sqlDB.SetMaxIdleConns(10)           // 最大空闲连接
	sqlDB.SetMaxOpenConns(100)          // 最大连接数
	sqlDB.SetConnMaxLifetime(time.Hour) // 连接最大存活时间
	sqlDB.SetConnMaxIdleTime(30 * time.Minute)
	if err := db.AutoMigrate(&model.SimpleWeather{}); err != nil {
		return err
	}
	return nil

}

func Getdb() *gorm.DB {
	return db
}
