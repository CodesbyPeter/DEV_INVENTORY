package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"DEV_INVENTORY/internal/handlers"
	"DEV_INVENTORY/internal/logics"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	// -- Initial guide and first action --
	// Calling Guide once to get the initial option
	initialOption := logics.Guide(scanner)
	handleOption(initialOption, scanner)
	
	// -- Continous loop for subsequent actions --
	// Infinite loop that breaks on Exit
	for {
		fmt.Println("\n Enter your next option: ")
		fmt.Print("Type here (1 for Learning, 2 for Project, 3 to exit): ")
		if !scanner.Scan() {
			break // Exit if scanning fails
		}
		input := strings.TrimSpace(scanner.Text())
		// Checking for Exit command first
		if strings.EqualFold(input, "exit") || input == "3" {
			fmt.Println("Exciting Dev_Inventory. GoodBye.")
			return
		}
		// Trying passing input as an interger
		nextOption, err := strconv.Atoi(input) // strconv.Atoi used for robust integer conversion
		if err != nil {
			fmt.Println("Invalid input. Please enter a number (1,2 0r 3) or 'exit'.")
			continue
		}
		handleOption(nextOption, scanner) // Helper function to be used for subsequent functions
	}
}

func handleOption(option int, scanner *bufio.Scanner) {
	switch option {
	case 1:
		handlers.HandleLearning(scanner)
	case 2:
		handlers.HandleProject(scanner)
	case 3:
		fmt.Println("Exciting Dev_Inventory. GoodBye.")
		os.Exit(0)
	default:
		fmt.Println("Invalid selection. Please choose 1 for learning, 2 for project or 3 to exit.")
	}
}
