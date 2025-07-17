package handlers

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func HandleRequests() {
	// HandleLearning() // Calls the Learning function
	// HandleProject() // Calls the Project function
}

// HandleLearning prompts the user for their specific learning choice.
// It now takes a scanner to read subsequent input.
func HandleLearning(scanner *bufio.Scanner) {
	fmt.Println("Are you: \n 1. Learning something new \n 2. Continuing with something you have learned before")
	fmt.Print("Type here: ")
	if !scanner.Scan() {
		fmt.Println("Errror reading your Input.")
		return
	}
	input := strings.TrimSpace(scanner.Text())
	option, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid input. Please enter 1 or 2.")
		return
	}
	switch option {
	case 1:
		fmt.Println("Great! Start journaling your new learning journey.")
	case 2:
		fmt.Println("Excellent! Keep pushing forward with your continuous learning.")
	default:
		fmt.Println("Invalid learning option. Please choose 1 or 2.")
	}
}

// / HandleProject prompts the user for their specific project choice.
// It now takes a scanner to read subsequent input.
func HandleProject(scanner *bufio.Scanner) {
	fmt.Println("Are you: \n 1. Working on a new project \n 2. Continuing with an existing one")
	fmt.Print("Type here: ")
	if !scanner.Scan() {
		fmt.Println("Errror reading your Input.")
		return
	}
	input := strings.TrimSpace(scanner.Text())
	option, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid input. Please enter 1 or 2.")
		return
	}
	switch option {
	case 1:
		fmt.Println("Great! What have you started building? ")
	case 2:
		fmt.Println("Excellent! What feature have you implemented today for your project? ")
	default:
		fmt.Println("Invalid Project option. Please choose 1 or 2.")
	}
}

// DefaultSelection informs the user that no valid option was chosen.
func DefaultSelection() {
	fmt.Println("Invalid selection. Please choose 1 for Learning or 2 for Project.")
}
