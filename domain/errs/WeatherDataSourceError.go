package errs

import (
	"errors"
	"fmt"
)

const (
	WeatherDataSourceErrorUnKnown  = 500
	WeatherDataSourceErrorExternal = 500
	WeatherDataSourceErrorRequest  = 400
)

func WeatherDataSourceErrorType(err error) int {
	var target WeatherDataSourceError
	if errors.As(err, &target) {
		return target.ErrType
	}
	return WeatherDataSourceErrorUnKnown
}

type WeatherDataSourceError struct {
	OriginErr error
	ErrType   int
}

func IsWeatherDataSourceError(err error) bool {
	var target WeatherDataSourceError
	return errors.As(err, &target)
}

func (s WeatherDataSourceError) Error() string {
	return fmt.Errorf("streamer error: %w", s.OriginErr).Error()
}

func (s WeatherDataSourceError) UnWrap() error {
	return s.OriginErr
}
