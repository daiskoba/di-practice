package usecase

import "fmt"

type WeatherForecast interface {
	UsecaseForecast(place string)
}

type WeatherClient interface {
	Forecast(place string) (*Weather, error)
}

type weatherForecast struct {
	client WeatherClient
}

type Weather struct {
	PublicTime string
	Summary    string
	Forecasts  map[string]string
}

func NewForecast(client WeatherClient) WeatherForecast {

	return &weatherForecast{
		client: client,
	}
}

func (wf *weatherForecast) UsecaseForecast(place string) {
	weather, _ := wf.client.Forecast(place)

	fmt.Printf("Weather summmary %v\n", weather.Summary)
	fmt.Printf("Weather forcasts...\n")
	for k, v := range weather.Forecasts {
		fmt.Printf("\t%v %v\n", k, v)
	}
	fmt.Printf("Update: %v", weather.PublicTime)
}
