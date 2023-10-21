package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/erickeniuk/geography"
)

type Weather struct {
	Latitude             float64     `json:"latitude"`
	Longitude            float64     `json:"longitude"`
	GenerationtimeMs     float64     `json:"generationtime_ms"`
	UtcOffsetSeconds     int         `json:"utc_offset_seconds"`
	Timezone             string      `json:"timezone"`
	TimezoneAbbreviation string      `json:"timezone_abbreviation"`
	Elevation            float64     `json:"elevation"`
	HourlyUnits          HourlyUnits `json:"hourly_units"`
	Hourly               Hourly      `json:"hourly"`
}
type HourlyUnits struct {
	Time                string `json:"time"`
	Temperature2M       string `json:"temperature_2m"`
	ApparentTemperature string `json:"apparent_temperature"`
	Precipitation       string `json:"precipitation"`
	Cloudcover          string `json:"cloudcover"`
	Windspeed10M        string `json:"windspeed_10m"`
	Winddirection10M    string `json:"winddirection_10m"`
}
type Hourly struct {
	Time                []int64   `json:"time"`
	Temperature2M       []float64 `json:"temperature_2m"`
	ApparentTemperature []float64 `json:"apparent_temperature"`
	Precipitation       []float64 `json:"precipitation"`
	Cloudcover          []int     `json:"cloudcover"`
	Windspeed10M        []float64 `json:"windspeed_10m"`
	Winddirection10M    []int     `json:"winddirection_10m"`
}

//type WeatherData struct

type WeatherDisplayHourly struct {
	City      string
	Forecasts []ForecastHourly
}

type ForecastHourly struct {
	Date        string
	Temperature string
}

func GetWeather(latLong geography.LatLong) (*Weather, error) {
	endpoint := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%.6f&longitude=%.6f&hourly=temperature_2m,apparent_temperature,precipitation,cloudcover,windspeed_10m,winddirection_10m&temperature_unit=fahrenheit&windspeed_unit=mph&precipitation_unit=inch&forecast_days=1&timeformat=unixtime", latLong.Latitude, latLong.Longitude)
	resp, err := http.Get(endpoint)
	if err != nil {
		return nil, fmt.Errorf("error making request to Weather API: %w", err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body) // response body is []byte

	var response Weather
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("Error while trying to decode response :( \n%w", err)
	}

	return &response, nil
}

func ExtractHourlyWeatherData(city string, rawWeather Weather) (WeatherDisplayHourly, error) {

	var forecasts []ForecastHourly

	for i, unix_time := range rawWeather.Hourly.Time {
		fmt.Println("rawWeather.Hourly.Time i: ", i)
		fmt.Println("rawWeather.Hourly.Time unix_time: ", unix_time)
		date := time.Unix(unix_time, 0)
		fmt.Println("converted date: ", date)
		forecast := ForecastHourly{
			Date:        date.Format("Mon 15:04"),
			Temperature: fmt.Sprintf("%.1f°F", rawWeather.Hourly.Temperature2M[i]),
		}
		forecasts = append(forecasts, forecast)
		fmt.Println("forecasts: ", forecasts)
	}

	return WeatherDisplayHourly{
		City:      city,
		Forecasts: forecasts,
	}, nil

}
