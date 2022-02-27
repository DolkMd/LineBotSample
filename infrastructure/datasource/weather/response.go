package infraweather

type WeatherValue struct {
	ID          float64 `json:"id"`
	Main        string  `json:"main"`
	Description string  `json:"description"`
	Icon        string  `json:"icon"`
}
type CurrentWeather struct {
	Dt           int64          `json:"dt"`
	Sunrise      float64        `json:"sunrise"`
	Sunset       float64        `json:"sunset"`
	Temp         float64        `json:"temp"`
	FeelsLike    float64        `json:"feels_like"`
	Pressure     float64        `json:"pressure"`
	Humidity     float64        `json:"humidity"`
	DewPofloat64 float64        `json:"dew_pofloat64"`
	Uvi          float64        `json:"uvi"`
	Clouds       float64        `json:"clouds"`
	Visibility   float64        `json:"visibility"`
	WindSpeed    float64        `json:"wind_speed"`
	WindDeg      float64        `json:"wind_deg"`
	WindGust     float64        `json:"wind_gust"`
	Weather      []WeatherValue `json:"weather"`
}
type MinutelyWeather struct {
	Dt            float64 `json:"dt"`
	Precipitation float64 `json:"precipitation"`
}
type HourlyWeather struct {
	Dt           int64          `json:"dt"`
	Temp         float64        `json:"temp"`
	FeelsLike    float64        `json:"feels_like"`
	Pressure     float64        `json:"pressure"`
	Humidity     float64        `json:"humidity"`
	DewPofloat64 float64        `json:"dew_pofloat64"`
	Uvi          float64        `json:"uvi"`
	Clouds       float64        `json:"clouds"`
	Visibility   float64        `json:"visibility"`
	WindSpeed    float64        `json:"wind_speed"`
	WindDeg      float64        `json:"wind_deg"`
	WindGust     float64        `json:"wind_gust"`
	Weather      []WeatherValue `json:"weather"`
	Pop          float64        `json:"pop"`
}
type DailyWeatherTemp struct {
	Day   float64 `json:"day"`
	Min   float64 `json:"min"`
	Max   float64 `json:"max"`
	Night float64 `json:"night"`
	Eve   float64 `json:"eve"`
	Morn  float64 `json:"morn"`
}
type DailyWeatherFeelsLike struct {
	Day   float64 `json:"day"`
	Night float64 `json:"night"`
	Eve   float64 `json:"eve"`
	Morn  float64 `json:"morn"`
}
type DailyWeather struct {
	Dt           int64                 `json:"dt"`
	Sunrise      float64               `json:"sunrise"`
	Sunset       float64               `json:"sunset"`
	Moonrise     float64               `json:"moonrise"`
	Moonset      float64               `json:"moonset"`
	MoonPhase    float64               `json:"moon_phase"`
	Temp         DailyWeatherTemp      `json:"temp"`
	FeelsLike    DailyWeatherFeelsLike `json:"feels_like"`
	Pressure     float64               `json:"pressure"`
	Humidity     float64               `json:"humidity"`
	DewPofloat64 float64               `json:"dew_pofloat64"`
	WindSpeed    float64               `json:"wind_speed"`
	WindDeg      float64               `json:"wind_deg"`
	WindGust     float64               `json:"wind_gust"`
	Weather      []WeatherValue        `json:"weather"`
	Clouds       float64               `json:"clouds"`
	Pop          float64               `json:"pop"`
	Uvi          float64               `json:"uvi"`
	Rain         float64               `json:"rain,omitempty"`
}

type Weather struct {
	Lat            float64           `json:"lat"`
	Lon            float64           `json:"lon"`
	Timezone       string            `json:"timezone"`
	TimezoneOffset float64           `json:"timezone_offset"`
	Current        CurrentWeather    `json:"current"`
	Minutely       []MinutelyWeather `json:"minutely"`
	Hourly         []HourlyWeather   `json:"hourly"`
	Daily          []DailyWeather    `json:"daily"`
}
