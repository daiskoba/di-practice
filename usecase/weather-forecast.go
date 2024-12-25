package usecase

import "fmt"

type WeatherForecast interface {
	UsecaseForecast(place string)
}

type WheatherClient interface {
	Forecast(place string) (*Wheather, error)
}

type wheatherForecast struct {
	client WheatherClient
}

type Wheather struct {
	PublicTime string
	Summary    string
	Forecasts  map[string]string
}

func NewForecast(client WheatherClient) WeatherForecast {

	return &wheatherForecast{
		client: client,
	}
}

func (wf *wheatherForecast) UsecaseForecast(place string) {
	wheather, _ := wf.client.Forecast(place)

	fmt.Printf("Wheather summmary %v\n", wheather.Summary)
	fmt.Printf("Wheather forcasts...\n")
	for k, v := range wheather.Forecasts {
		fmt.Printf("\t%v %v\n", k, v)
	}
	fmt.Printf("Update: %v", wheather.PublicTime)
}
