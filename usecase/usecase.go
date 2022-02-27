package usecase

import (
	"context"
	"net/http"
	"strings"

	"github.com/DolkMd/LineBotSample/domain/interfaces"
	"github.com/DolkMd/LineBotSample/domain/interfaces/model"
)

type UseCaseConfig struct {
	Longitude, Latitude float64
}
type UseCase interface {
	PostWeather(ctx context.Context) error
	HandleStreamerCallback(ctx context.Context, req *http.Request) error
	NotifyError(ctx context.Context, message string) error
}
type ucase struct {
	streamer            interfaces.Streamer
	weatherDsrc         interfaces.WeatherDataSource
	longitude, latitude float64
}

func NewUseCase(streamer interfaces.Streamer, weatherDsrc interfaces.WeatherDataSource, conf UseCaseConfig) UseCase {
	return &ucase{
		streamer:    streamer,
		weatherDsrc: weatherDsrc,
		longitude:   conf.Longitude,
		latitude:    conf.Latitude,
	}
}

func (u *ucase) PostWeather(ctx context.Context) error {
	from, to, err := u.weatherDsrc.GetWeather1HoursRange(u.longitude, u.latitude)
	if err != nil {
		return err
	}

	if from.Value == model.Clear && to.Value == model.Bad {
		return u.streamer.SendText(ctx, "天気荒れるで: "+to.Time.String())
	}

	return nil
}

func (u *ucase) NotifyError(ctx context.Context, message string) error {
	return u.streamer.SendText(ctx, message)
}

func (u *ucase) HandleStreamerCallback(ctx context.Context, req *http.Request) error {
	return u.streamer.HandleCallback(ctx, req, func(cmdTextBox string) (string, error) {
		cmds := strings.Split(cmdTextBox, " ")
		if len(cmds) > 0 {
			switch cmds[0] {
			case "weather":
				from, to, err := u.weatherDsrc.GetWeather1HoursRange(u.longitude, u.latitude)
				if err != nil {
					return "", err
				}
				return from.Time.String() + "/" + string(from.Value) + "\n" + to.Time.String() + "/" + string(to.Value), nil
			}
		}
		return cmdTextBox, nil
	})
}
