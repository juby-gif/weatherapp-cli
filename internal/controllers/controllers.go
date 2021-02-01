package controllers

import (
	"fmt"
	"time"

	"github.com/juby-gif/weatherapp-cli/internal/utils"
)

func ProcessRecordCreation() {
	addRecord()
	time.Sleep(5 * time.Second)
}

func ProcessRecordListing() {

}

func ProcessCalculation() {

}

func addRecord() {
	utils.GetScreenCleared()
	fmt.Println("###############################")
	fmt.Println("# Weather App Command Console #")
	fmt.Println("###############################")
	fmt.Printf("\n")
	fmt.Println("Please enter the following records:")
	fmt.Printf("\n")
	fmt.Println("Temperature (deg):")
	temperature := utils.ReadConsoleFloat64()
	fmt.Println("Humidity (%):")
	humidity := utils.ReadConsoleFloat64()
	fmt.Println("Pressure (Pa):")
	pressure := utils.ReadConsoleFloat64()
	fmt.Println("CO2 (ppm):")
	co2 := utils.ReadConsoleFloat64()
	fmt.Println("TVOC (ppb):")
	tvoc := utils.ReadConsoleFloat64()
	fmt.Println("Date/Time:")
	date := utils.ReadConsoleFloat64()
	fmt.Printf("\n")

	fmt.Println(temperature, humidity, pressure, co2, tvoc, date)

	// time_series_datum := &TimeSeriesDatum{
	// 	temperature: temperature,
	// 	humidity:    humidity,
	// 	pressure:    pressure,
	// 	co2:         co2,
	// 	tvoc:        tvoc,
	// 	date:        date,
	// }
	// time_series_datum = append(time_series_datum, tsd)
}
