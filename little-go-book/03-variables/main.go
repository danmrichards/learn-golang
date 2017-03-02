package main

import (
	"fmt"
)

func main() {
	// Short hand for this (power := 9000) will infer the type.
	var power int
	power = 9000

	fmt.Printf("It's over %d\n", power)
}
