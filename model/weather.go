package model

import "time"

type Weather struct {
	ID   uint   `gorm:"primaryKey"`
	City string `gorm:"index"`
	Date string `gorm:"index"`

	MinTemp float64
	MaxTemp float64
	Sky     string

	Humidity  float64
	Rain      float64
	WindSpeed float64
	AQI       float64

	CreatedAt time.Time
}