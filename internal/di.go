package di

import (
	"github-sub/daiskoba/di-practice/infra/wheather"
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
	uf := usecase.NewForecast(c.WheatherClient())
	uf.UsecaseForecast(place)
}

// Infra

func (c *Container) WheatherClient() usecase.WheatherClient {
	return wheather.NewWheatherClient()
}
