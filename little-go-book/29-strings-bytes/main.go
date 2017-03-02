package main

import (
	"fmt"
)

func main() {
	stringA := "the spice must flow"
	fmt.Printf("This is a string: %v\n", stringA)
	fmt.Println("This is what happens when we iterate:")
	for _, value := range stringA {
		fmt.Printf("%c\n", value)
	}

	bytes := []byte(stringA)
	fmt.Printf("This is a byte array: %v\n", bytes)
	fmt.Println("This is what happens when we iterate:")
	for _, value := range bytes {
		fmt.Printf("%v\n", value)
	}

	stringB := string(bytes)
	fmt.Printf("This is a string again: %v\n", stringB)
	fmt.Println("This is what happens when we iterate:")
	for _, value := range stringB {
		fmt.Printf("%v\n", value)
	}
}
