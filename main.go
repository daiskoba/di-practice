package main

import (
	"github-sub/daiskoba/di-practice/infra/wheather"
	"github-sub/daiskoba/di-practice/usecase"
)

func main() {
	client := wheather.NewWheatherClient()
	uf := usecase.NewForecast(client)

	place := "tokyo"
	uf.UsecaseForecast(place)
}
