package main

import (
    "fmt"
)

// Person structure.
type Person struct {
    Name string
}

// Introduce yourself.
func (p *Person) Introduce() {
    fmt.Printf("Hi, I'm %s\n", p.Name)
}

// Saiyan structure.
// Go supports composition, which is including one structure into another
type Saiyan struct {
    *Person
    Power int
}

func main() {
    goku := &Saiyan{
        Person: &Person{"Goku"},
        Power: 9001,
    }

    goku.Introduce()
}
