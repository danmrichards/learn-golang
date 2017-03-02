package main

import (
	"fmt"
)

// Maps in Go are what other languages call hashtables or dictionaries. They
// work as you expect: you define a key and value, and can get, set and
// delete values from it.
func main() {
	lookup := make(map[string]int)

	// Maps grow dynamically. However, we can supply a second argument to make
	// to set an initial size. If you have some idea of how many keys your map
	// will have, defining an initial size can help with performance.
	// lookup := make(map[string]int, 100)

	// Assign a value.
	lookup["goku"] = 9001

	// A shorthand way to create and assign a slice is similar to arrays.
	// lookup := map[string]int{
	//    "goku": 9001,
	//    "gohan": 2044,
	// }

	// Prints 0, false
	// 0 is the default value for an integer
	power, exists := lookup["vegeta"]
	fmt.Println(power, exists)

	// Returns 1
	total := len(lookup)
	fmt.Printf("The map has a length of %d\n", total)

	// Has no return, can be called on a non-existing key
	delete(lookup, "goku")
}
