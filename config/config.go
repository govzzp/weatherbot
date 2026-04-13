package config

import (
	"os"

	"github.com/goccy/go-yaml"
)

type Config struct {
	Feishu struct {
		Webhook string
	}
	Caiyun struct {
		Token string
	}
	Cities []struct {
		Name string
		Lng  float64
		Lat  float64
	}
	Mysql struct {
		DSN string
	}
	Server struct {
		Port int
	}
}

func LoadConfig() *Config {
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}

	var cfg Config
	_ = yaml.Unmarshal(data, &cfg)

	return &cfg
}
