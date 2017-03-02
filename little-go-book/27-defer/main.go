package main

import (
	"fmt"
)

func main() {
	// Defer informs Go that the particular function should be run once the
	// enclosing function has returned. This is useful if we have large
	// functions with multiple return points or a resource we need to close.
	defer done()

	for i := 1; i < 20; i++ {
		fmt.Printf("At number %d\n", i)
	}
}

func done() {
	fmt.Println("Done now!")
}
