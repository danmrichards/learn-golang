package main

import (
	"fmt"
)

// Functions are first-class types which can then be used anywhere: as a field
// type, as a parameter, as a return value

// Add - Adds 2 integers together.
type Add func(a int, b int) int

func main() {
	// Using functions like this can help decouple code from specific
	// implementations much like we achieve with interfaces.
	fmt.Println(process(func(a int, b int) int {
		return a + b
	}))
}

func process(adder Add) int {
	return adder(1, 2)
}
