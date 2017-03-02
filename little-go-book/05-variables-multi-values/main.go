package main

import (
	"fmt"
)

func main() {
	power := 1000
	fmt.Printf("default power is %d\n", power)

	// Multiple variables can be created and assigned at the same time.
	// Notice that the presence of a new variable (name) allows power to be
	// assigned a new value!
	name, power := "Goku", 9000
	fmt.Printf("%s's power is over %d\n", name, power)

	// If we create a variable but don't use it, we'll get a compiler error.
	// newName, newPower := 1000
	// fmt.Printf("new default power is %d\n", newPower)
}
