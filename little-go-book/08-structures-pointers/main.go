package main

import (
    "fmt"
)

// Saiyan structure.
type Saiyan struct {
    Name string
    Power int
}

func main() {
    // Note: We're using the & operator to get the 'address' of the created
    // Saiyan. This ensures the same instance of the structure is modified
    // when we pass to Super. If we don't use the 'address' a copy of the
    // instance gets created and that is updated instead.
    goku := &Saiyan{"Goku", 9000}

    Super(goku)

    // Expected value = 19000
    fmt.Println(goku.Power)

    SuperAgain(goku)

    // Expected value = 19000
    // We've changed the address in SuperAgain, but goku still stats.
    fmt.Println(goku.Power)
}

// Super Saiyan!
// Takes a pointer to an address of a Saiyan structure.
func Super(s *Saiyan) {
    s.Power += 10000
}

// SuperAgain - Demonostrates what happens when you try to change the address
// of the provided pointer.
// Takes a pointer to an address of a Saiyan structure.
func SuperAgain(s *Saiyan) {
    s = &Saiyan{"Gohan", 1000}
}
