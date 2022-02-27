package infraweather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/DolkMd/LineBotSample/domain/errs"
	"github.com/DolkMd/LineBotSample/domain/interfaces"
	"github.com/DolkMd/LineBotSample/domain/interfaces/model"
)

type weather struct {
	apiKey string
}

func NewWeatherDataSource(apiKey string) interfaces.WeatherDataSource {
	return &weather{
		apiKey: apiKey,
	}
}

func (w *weather) GetWeather1HoursRange(longitude, latitude float64) (model.Weather, model.Weather, error) {
	url := fmt.Sprintf(
		"https://api.openweathermap.org/data/2.5/onecall?lat=%f&lon=%f&units=metric&lang=ja&appid=%s",
		latitude,
		longitude,
		w.apiKey,
	)
	response, err := http.Get(url)
	if err != nil {
		return model.Weather{}, model.Weather{}, errs.WeatherDataSourceError{OriginErr: err, ErrType: response.StatusCode}
	}
	defer response.Body.Close()

	byts, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return model.Weather{}, model.Weather{}, errs.WeatherDataSourceError{OriginErr: err}
	}

	var weather Weather
	if err := json.Unmarshal(byts, &weather); err != nil {
		return model.Weather{}, model.Weather{}, errs.WeatherDataSourceError{OriginErr: err}
	}

	hourWeather := HourlyWeather{}
	for _, hour := range weather.Hourly {
		if weather.Current.Dt < hour.Dt {
			hourWeather = hour
			break
		}
	}

	nowWeather := model.Weather{Value: model.Clear, Time: time.Unix(weather.Current.Dt, 0)}
	if isBadWeather(weather.Current.Weather[0].Main) {
		nowWeather.Value = model.Bad
	}

	nextWeather := model.Weather{Value: model.Clear, Time: time.Unix(hourWeather.Dt, 0)}
	if isBadWeather(hourWeather.Weather[0].Main) {
		nextWeather.Value = model.Bad
	}

	return nowWeather, nextWeather, nil
}

func isBadWeather(weather string) bool {
	weather = strings.ToLower(weather)
	return weather == "rain" ||
		weather == "snow" ||
		weather == "extreme"

}
