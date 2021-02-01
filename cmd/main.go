package main

import (
	"fmt"
	_ "weatherapp-cli/utils"
)

func main() {
	// for {
	// 	choice := utils.ReadConsoleString()
	// 	if choice == "a" {
	menu()
	// 	} else {
	// 		break
	// 	}
	// }
}

func menu() {
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
