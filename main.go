package main

import (
	"fmt"
)

func initialModel() []string {
	choices := []string{
		"1) Add contact",
		"2) List contacts",
		"3) Remove contact",
		"4) Update contact",
		"5) Quit",
	}
	return choices
}
func main() {
	for i := 0; i < len(initialModel()); i++ {
		fmt.Print(initialModel()[i] + "\n")

	}
	for {

	}
}
