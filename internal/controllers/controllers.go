package controllers

import (
	"fmt"
	"time"

	"github.com/juby-gif/weatherapp-cli/internal/models"
	"github.com/juby-gif/weatherapp-cli/internal/utils"
)

func AddRecord() models.TimeSeriesDatum {
	utils.GetScreenCleared()
	fmt.Println("###############################")
	fmt.Println("# Weather App Command Console #")
	fmt.Println("###############################")
	fmt.Printf("\n")
	fmt.Println("Please enter the following records:")
	fmt.Printf("\n")
	fmt.Println("Temperature (deg):")
	temp := utils.ReadConsoleFloat64()
	fmt.Println("Humidity (%):")
	humidity := utils.ReadConsoleFloat64()
	fmt.Println("Pressure (Pa):")
	pressure := utils.ReadConsoleFloat64()
	fmt.Println("CO2 (ppm):")
	co2 := utils.ReadConsoleFloat64()
	fmt.Println("TVOC (ppb):")
	tvoc := utils.ReadConsoleFloat64()
	fmt.Println("Date/Time:")
	timestamp := time.Now()
	fmt.Printf("\n")
	utils.GetScreenCleared()

	time_series_datum := models.TimeSeriesDatum{
		Temperature: temp,
		Humidity:    humidity,
		Pressure:    pressure,
		Co2:         co2,
		Tvoc:        tvoc,
		Timestamp:   timestamp,
	}
	return time_series_datum
}

func ProcessRecordListing(tsd []models.TimeSeriesDatum) {
	fmt.Println("###############################")
	fmt.Println("# Weather App Command Console #")
	fmt.Println("###############################")
	for _, v := range tsd {
		fmt.Println("")
		fmt.Printf("T:%f \n", v.Temperature)
		fmt.Printf("H:%f \n", v.Humidity)
		fmt.Printf("P:%f \n", v.Co2)
		fmt.Printf("C:%f \n", v.Co2)
		fmt.Printf("T:%f \n", v.Tvoc)
		fmt.Println("Time:", v.Timestamp)
		fmt.Println("")
	}
}

func ProcessCalculation() {

}
