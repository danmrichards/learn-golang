package main

import (
	"fmt"
)

func main() {
	power := getPower()

	// If we try to create and assign (:=) the same variable again, we'll get
	// a compiler error - "no new variables on left side of :=".
	// power := 9000

	fmt.Printf("It's over %d\n", power)
}

func getPower() int {
	return 9001
}
