// FILE: weatherapp-cli/internal/models/models.go
package models

type TimeSeriesDatum struct {
	temperature float64
	humidity    float64
	pressure    float64
	co2         float64
	tvoc        float64
	date        string
}

type TimeSeriesData struct {
	tsd []*TimeSeriesDatum
}
