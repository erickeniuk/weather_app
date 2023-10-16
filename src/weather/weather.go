package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/erickeniuk/geography"
)

type Weather struct {
	Latitude             float64    `json:"latitude"`
	Longitude            float64    `json:"longitude"`
	GenerationtimeMs     float64    `json:"generationtime_ms"`
	UtcOffsetSeconds     int        `json:"utc_offset_seconds"`
	Timezone             string     `json:"timezone"`
	TimezoneAbbreviation string     `json:"timezone_abbreviation"`
	Elevation            float64    `json:"elevation"`
	DailyUnits           DailyUnits `json:"daily_units"`
	Daily                Daily      `json:"daily"`
}
type DailyUnits struct {
	Time               string `json:"time"`
	Temperature2MMax   string `json:"temperature_2m_max"`
	Temperature2MMin   string `json:"temperature_2m_min"`
	PrecipitationHours string `json:"precipitation_hours"`
}
type Daily struct {
	Time               []string  `json:"time"`
	Temperature2MMax   []float64 `json:"temperature_2m_max"`
	Temperature2MMin   []float64 `json:"temperature_2m_min"`
	PrecipitationHours []float64 `json:"precipitation_hours"`
}

func Get_Weather(latLong geography.LatLong) (*Weather, error) {
	endpoint := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%.6f&longitude=%.6f&timezone=auto&daily=temperature_2m_max,temperature_2m_min,precipitation_hours&temperature_unit=fahrenheit", latLong.Latitude, latLong.Longitude)
	resp, err := http.Get(endpoint)
	if err != nil {
		return nil, fmt.Errorf("error making request to Weather API: %w", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body) // response body is []byte
	fmt.Println()

	var response Weather
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("Error while trying to decode response :( \n%w", err)
	}

	return &response, nil
}
