package main

import (
	"bufio"
	"fmt"
	"os"
)

var contacts map[string]string

func printItems() []string {
	choices := []string{
		"1) Add contact",
		"2) List contacts",
		"3) Remove contact",
		"4) Update contact",
		"5) Quit",
	}
	for i := 0; i < len(choices); i++ {
		fmt.Print(choices[i] + "\n")

	}
	return choices
}

func miniCRM() {
	printItems()
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	// Step 1 : if input not integer between 1 and 5, ask again
	for {
		if input == "1\n" || input == "2\n" || input == "3\n" || input == "4\n" || input == "5\n" {
			break
		}
		fmt.Print("Invalid input. Please enter a number between 1 and 5: ")
		input, _ = reader.ReadString('\n')
	}

	if input == "1\n" {
		// prefix string "Enter contact name: "
		fmt.Print("Enter contact name: ")
		input, _ := reader.ReadString('\n')
		fmt.Print("name :", input)
		miniCRM()
	}
	// Function to make program persistent
	for {

	}
}

func main() {
	miniCRM()
}
