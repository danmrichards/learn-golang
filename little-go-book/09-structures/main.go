package main

import (
    "fmt"
)

// Saiyan structure.
type Saiyan struct {
    Name string
    Power int
}

// Super Saiyan!
// In this case we're setting the *Saiyan pointer as the receiver of the func.
func (s *Saiyan) Super() {
    s.Power += 10000
}

func main() {
    goku := &Saiyan{"Goku", 9000}

    // Call the function on the Saiyan address itself.
    goku.Super()

    // Expected value = 19000
    fmt.Println(goku.Power)
}
