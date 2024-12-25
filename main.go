package main

import (
	di "github-sub/daiskoba/di-practice/internal"
)

func main() {

	place := "tokyo"
	uf := di.NewContainer()
	uf.UsecaseForecast(place)
}
