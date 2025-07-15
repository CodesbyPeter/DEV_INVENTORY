package logics

import (
	"fmt"
)

func Guide() {
	// Welcome user and inquire about their next step
	fmt.Println("Welcome to Dev_Inventory......... \n Are you currently: \n 1. Learning \n 2. Building a Project? ")
}

func Guide_Choice(choice string) {
	switch choice {
	case "1":
		fmt.Println("If you are learning are you learning: \n 1. A new thing you have never journalled. \n 2. A continous thing you have ever journelled before. ")
	case "2":
		fmt.Println("Are you engaging in: \n 1. A new project \n 2. A Continuing/Existing one ")
	default:
		fmt.Println("Not Building or Learning. Relax there ain't much of that in the job market.")
	}
}
