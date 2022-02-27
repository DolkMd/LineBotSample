package interfaces

import (
	"context"
	"net/http"

	"github.com/DolkMd/LineBotSample/domain/interfaces/model"
)

type (
	Application interface {
		Run(ctx context.Context, port int) error
	}
	WeatherDataSource interface {
		GetWeather1HoursRange(longitude, latitude float64) (model.Weather, model.Weather, error)
	}
	Streamer interface {
		SendText(ctx context.Context, text string) error
		HandleCallback(ctx context.Context, request *http.Request, cmdCallback func(string) (string, error)) error
	}
)
