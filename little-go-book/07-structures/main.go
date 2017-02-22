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
    // Simplest way to create a new value of the structure.
    // goku := Saiyan{
    //     Name: "Goku",
    //     Power: 9000,
    // }

    // Empty value.
    // goku := Saiyan{}

    // Set parameters after creation.
    // goku := Saiyan{Name: "Goku"}
    // goku.Power = 9000

    // The field names can be skipped too (only recommended for simple cases).
    goku := Saiyan{"Goku", 9000}
    fmt.Println(goku.Power)
}
