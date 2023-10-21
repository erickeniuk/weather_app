package geography

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

type GeoResponse struct {
	// A list of results; we only need the first one
	Results []LatLong `json:"results"`
}

type LatLong struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func BeginMessage() string {
	// Return a basic message that the app is beginning from the georgraphy module
	begin_message := "beginning"
	message := fmt.Sprintf("I am ... %s !\n", begin_message)
	return message
}

func GetLatLong(city string) (*LatLong, error) {
	endpoint := fmt.Sprintf("https://geocoding-api.open-meteo.com/v1/search?name=%s&count=1&language=en&format=json", url.QueryEscape(city))
	resp, err := http.Get(endpoint)
	if err != nil {
		fmt.Printf("%s", err)
		return nil, fmt.Errorf("Error while trying to get lat long via Geo API: %w", err)
	}
	defer resp.Body.Close()

	var response GeoResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("Error while trying to decode response: %w", err)
	}

	if len(response.Results) < 1 {
		println("city provided: [%s]", city)
		return nil, errors.New("No results found for city provided.")
	}
	return &response.Results[0], nil
}
