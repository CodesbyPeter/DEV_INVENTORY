package logics

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func Guide(scanner *bufio.Scanner) int {
	// Welcome user and inquire about their next step
	fmt.Println("Welcome to Dev_Inventory......... \n Are you currently: \n 1. Learning \n 2. Building a Project? ")
	fmt.Print("Type here: ")

	if !scanner.Scan() {
		fmt.Println("Error reading initial input. Exiting.")
		return 0
	}

	input := strings.TrimSpace(scanner.Text()) // Get text from scanner and trim
	option, err := strconv.Atoi(input)         // Convert to int using strconv.Atoi
	if err != nil {
		fmt.Println("Invalid initial input. Please enter a number (1, 2 or 3).")
		return 0
	}
	fmt.Println("Type '3' or 'EXIT' at any time to quit the program.")
	return option
}
