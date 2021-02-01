package main

import (
	"fmt"

	"time"

	"github.com/juby-gif/weatherapp-cli/internal/controllers"
	_ "github.com/juby-gif/weatherapp-cli/internal/models"
	"github.com/juby-gif/weatherapp-cli/internal/utils"
)

func main() {

	for {
		menu()
		choice := utils.ReadConsoleString()
		switch choice {
		case "a", "A":
			controllers.ProcessRecordCreation()

		case "b", "B":
			controllers.ProcessRecordListing()

		case "c", "C":
			controllers.ProcessCalculation()

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
	fmt.Println("###############################")
	fmt.Println("# Weather App Command Console #")
	fmt.Println("###############################")
	fmt.Printf("\n")
	fmt.Println("Menu")
	fmt.Println("(A) Add a record")
	fmt.Println("(B) List all records")
	fmt.Println("(C) Calculate Summary")
	fmt.Println("(D) Exit")
	fmt.Printf("\n")
	fmt.Println("Please enter choice:")
}
