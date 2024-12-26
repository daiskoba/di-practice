package di

import (
	"github-sub/daiskoba/di-practice/infra/weather"
	"github-sub/daiskoba/di-practice/usecase"
)

type Container struct {
	cache map[string]any
}

func NewContainer() *Container {
	return &Container{
		cache: map[string]any{},
	}
}

// Usecase
func (c *Container) UsecaseForecast(place string) {
	uf := usecase.NewForecast(c.WeatherClient())
	uf.UsecaseForecast(place)
}

// Infra

func (c *Container) WeatherClient() usecase.WeatherClient {
	return weather.NewWeatherClient()
}
