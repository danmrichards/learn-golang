package main

import (
	"fmt"
)

// In Go arrays are fixed. Declaring an array requires that we specify the
// size, and once the size is specified, it cannot grow.
var scores [4]int

func main() {
	// We can initialize the array with values.
	scores := [4]int{9001, 9333, 212, 33}

	// We can use len to get the length of the array.
	fmt.Printf("The array has length: %d\n", len(scores))

	// The range function can be used to iterate over it.
	for index, value := range scores {
		fmt.Printf("%d:%d\n", index, value)
	}
}
