package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"DEV_INVENTORY/internal/logics"
	//"DEV_INVENTORY/internal/learning"
	// "DEV_INVENTORY/internal/logics"
)

func main() {
	// for {
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

		switch strings.TrimSpace(input) {
		case "1":
			fmt.Println("If you are learning\n are you learning: \n 1. A new thing you have never journalled. \n 2. A continous thing you have ever journelled before. ")
		case "2":
			fmt.Println("Are you engaging in: \n 1. A new project \n 2. A Continuing/Existing one ")
		default:
			fmt.Println("Not Building or Learning. Relax there ain't much of that in the job market.")
		}
	}

	// learning.Learn_Choices(input[1])

	// }
	// Call guide
	// Call guide_choice
}

func choice_resolver() {
}
