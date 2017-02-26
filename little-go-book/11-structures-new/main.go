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
    // Despite the lack of constructors, Go does have a built-in new function
    // which is used to allocate the memory required by a type.
    // The result of new(X) is the same as &X{}.
    goku := new(Saiyan)
    goku.Name = "goku"
    goku.Power = 9001

    // Most people prefer this style whenever they have fields to initialize,
    // since it tends to be easier to read.
    // goku := &Saiyan {
    //     Name: "goku",
    //     Power: 9000,
    // }

    // Expected value = 9001
    fmt.Println(goku.Power)
}
