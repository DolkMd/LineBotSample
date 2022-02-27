package model

import "time"

type WeatherValue string

const (
	Clear WeatherValue = "clear"
	Bad   WeatherValue = "bad"
)

type Weather struct {
	Value WeatherValue
	Time  time.Time
}
