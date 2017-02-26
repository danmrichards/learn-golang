package main

import (
    "fmt"
)

// Saiyan structure.
type Saiyan struct {
    Name string
    Power int
}

// Structures don’t have constructors. Instead, you create a function that
// returns an instance of the desired type (like a factory).
// Note: This function is returning a pointer, but we don't have to.
func newSaiyan(name string, power int) *Saiyan {
    return &Saiyan{
        Name: name,
        Power: power,
    }
}

func main() {
    goku := newSaiyan("Goku", 9000)

    // Expected value = 9000
    fmt.Println(goku.Power)
}
