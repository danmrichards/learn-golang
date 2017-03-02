package main

import (
	"fmt"
)

func main() {
	count := 11

	// Go supports a slightly modified if-statement, one where a value can be
	// initiated prior to the condition being evaluated.
	if x := 10; count > x {
		fmt.Println("Yarp")
	}
}
