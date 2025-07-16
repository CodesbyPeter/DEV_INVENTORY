package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"DEV_INVENTORY/internal/logics"

)

func main() {
	logics.Guide()
	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input=strings.TrimSpace(input)
		if input == "EXIT" {
			break
		}
       // Handling Learning/Project directly in main file brings about issues in messiness and scaling
		/* switch strings.TrimSpace(input) {
		case "1":
			fmt.Println("If you are learning\n are you learning: \n 1. A new thing you have never journalled. \n 2. A continous thing you have ever journelled before. ")
		case "2":
			fmt.Println("Are you engaging in: \n 1. A new project \n 2. A Continuing/Existing one ")
		default:
			fmt.Println("Not Building or Learning. Relax there ain't much of that in the job market.")
		} */
	}
}

// Extracting logic from each case into their own functions

// Learning Function 
func HandleLearning() {
	fmt.Println("Are you: \n 1. Learning something new \n 2. Continuing with something you have learned before")
}

// Project Function
func HandleProject() {
	fmt.Println("Are you: \n 1. Working on a new project \n 2. Continuing with an existing one")
}

