package main

import (
	"context"
	"os"
	"strconv"

	"github.com/fcfcqloow/go-advance/log"
	"github.com/joho/godotenv"

	"github.com/DolkMd/LineBotSample/aop"
	_ "github.com/DolkMd/LineBotSample/aop"
	infragin "github.com/DolkMd/LineBotSample/infrastructure/application/gin"
	infraweather "github.com/DolkMd/LineBotSample/infrastructure/datasource/weather"
	"github.com/DolkMd/LineBotSample/infrastructure/streamer/infraline"
	"github.com/DolkMd/LineBotSample/usecase"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Warn("fail loading .env", err)
	}
	aop.InitLogger(os.Getenv("LOG_LEVEL"))
}

func main() {
	port := os.Getenv("PORT")
	groupID := os.Getenv("LINE_GROUP_ID")
	channelToken := os.Getenv("LINE_CHANNEL_TOKEN")
	channelSecret := os.Getenv("LINE_CHANNEL_SECRET")
	openWeatherApiKey := os.Getenv("OPEN_WEATHER_API_KEY")
	longitudeStr := os.Getenv("LONGITUDE")
	latitudeStr := os.Getenv("LATITUDE")

	portNum, err := strconv.Atoi(port)
	if err != nil {
		portNum = 8080
	}

	longitude, err := strconv.ParseFloat(longitudeStr, 64)
	if err != nil {
		panic(err)
	}

	latitude, err := strconv.ParseFloat(latitudeStr, 64)
	if err != nil {
		panic(err)
	}

	streamer, err := infraline.NewStreamer(channelToken, channelSecret, groupID)
	if err != nil {
		panic(err)
	}

	wDsrc := infraweather.NewWeatherDataSource(openWeatherApiKey)

	ucase := usecase.NewUseCase(streamer, wDsrc, usecase.UseCaseConfig{
		Longitude: longitude,
		Latitude:  latitude,
	})

	app := infragin.NewApplication(ucase)
	app.Run(context.Background(), portNum)
}
