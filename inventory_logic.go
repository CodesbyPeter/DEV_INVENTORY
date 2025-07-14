package main

import (
	"fmt"
)

func guide() {
	// Welcome user and inquire about their next step
	fmt.Println("Welcome to Dev_Inventory......... \n Are you currently: \n 1. Learning \n 2. Building a Project? ")
	// Storing the choice in a variable
	var choice string
	// Grabbing input from user so as to use it
	fmt.Scanln(&choice)
	// Switch statement to determine the next step
	switch {
	case choice == "1":
		fmt.Println("If you are learning what is it you are currently learning? ")
	case choice == "2":
		fmt.Println("Are you engaging in: \n 1. A new project \n 2. A Continuing/Existing one ")
	default:
		fmt.Println("Not Building or Learning. Relax there ain't much of that in the job market.")
	}
}
