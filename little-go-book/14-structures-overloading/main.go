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
// Go supports composition, which is including one structure into another.
type Saiyan struct {
    *Person
    Power int
}

// Introduce yourself, Saiyan!
// Simply, Go doesn’t support overloading. Because implicit composition is
// just a compiler trick, we can “overwrite” the functions of a composed type.
func (s *Saiyan) Introduce() {
    fmt.Printf("Hi, I'm %s. Ya!\n", s.Name)
}

func main() {
    goku := &Saiyan{
        Person: &Person{"Goku"},
        Power: 9001,
    }

    goku.Introduce()
}
