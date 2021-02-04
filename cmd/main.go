package main

import (
	"fmt"
	"time"

	"github.com/juby-gif/weatherapp-cli/internal/controllers"
	"github.com/juby-gif/weatherapp-cli/internal/models"
	"github.com/juby-gif/weatherapp-cli/internal/utils"
)

func main() {
	var tsd []models.TimeSeriesDatum
	fmt.Println(tsd)
	for {
		menu()
		choice := utils.ReadConsoleString()
		switch choice {
		case "a", "A":
			datum := controllers.AddRecord()
			tsd = append(tsd, datum)
			menuTopHeader()
			fmt.Println("Record created!")
			PressKey()

		case "b", "B":
			utils.GetScreenCleared()
			controllers.ProcessRecordListing(tsd)
			PressKey()

		case "c", "C":
			utils.GetScreenCleared()
			controllers.ProcessCalculation(tsd)
			PressKey()
		case "d", "D":
			utils.GetScreenCleared()
			menuTopHeader()
			fmt.Println("Have nice day!")
			PressKey()
			utils.GetScreenCleared()
		default:
			utils.GetScreenCleared()
			fmt.Println("Sorry the choice doesn't exist. Please try again!")
			time.Sleep(3 * time.Second)
		}
		utils.GetScreenCleared()
	}
}

func menu() {
	utils.GetScreenCleared()
	menuTopHeader()
	fmt.Println("Menu")
	fmt.Println("(A) Add a record")
	fmt.Println("(B) List all records")
	fmt.Println("(C) Calculate Summary")
	fmt.Println("(D) Exit")
	fmt.Printf("\n")
	fmt.Println("Please enter choice:")
}

func PressKey() {
	fmt.Println("")
	fmt.Println("PRESS ANY KEY TO CONTINUE")
	choice := utils.ReadConsoleString()
	switch choice {
	default:
		menu()
	}
}

func menuTopHeader() {
	fmt.Println("###############################")
	fmt.Println("# Weather App Command Console #")
	fmt.Println("###############################")
	fmt.Printf("\n")
}
