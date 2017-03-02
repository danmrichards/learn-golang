package main

import (
	"fmt"
)

// Saiyan structure.
// Fields can be of any type – including maps of other structures.
type Saiyan struct {
	Name    string
	Power   int
	Friends map[string]*Saiyan
}

func main() {
	goku := &Saiyan{
		Name:    "Goku",
		Power:   9001,
		Friends: make(map[string]*Saiyan),
	}

	goku.Friends["krillin"] = &Saiyan{
		Name:    "Krillin",
		Power:   1000,
		Friends: nil,
	}

	goku.Friends["dan"] = &Saiyan{
		Name:    "Dan",
		Power:   999999,
		Friends: nil,
	}

	fmt.Printf("%s has a power level of %d\n", goku.Name, goku.Power)

	for _, friend := range goku.Friends {
		fmt.Printf("His friend %s has a power level of %d\n", friend.Name, friend.Power)
	}
}
