package main

import (
	"fmt"
)

func main() {
	// If a function returns multiple values you can assign them to multiple
	// variables. If you don't care about one or more of them, assign to _.
	_, exists := power("Goku")

	// This will cause the exists values to be false
	// _, exists := power("Dan")

	if exists == false {
		fmt.Println("Person does not exist and has no power")
	}
}

// Gets the power level for a given person.
// It takes the name of the person as a string parameter.
// It returns the power level and a boolean indicating if the user exists.
func power(name string) (int, bool) {
	if name == "Goku" {
		return 9000, true
	}

	return 0, false
}
