package main

import (
	"fmt"
)

// Saiyan structure.
// Fields can be of any type – including other structures.
type Saiyan struct {
	Name   string
	Power  int
	Father *Saiyan
}

func main() {
	gohan := &Saiyan{
		Name:  "Gohan",
		Power: 1000,
		Father: &Saiyan{
			Name:   "Goku",
			Power:  9001,
			Father: nil,
		},
	}

	fmt.Println(gohan.Name)
	fmt.Println(gohan.Power)
	fmt.Println(gohan.Father.Name)
	fmt.Println(gohan.Father.Power)
}
