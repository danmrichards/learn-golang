package main

import (
    "fmt"
    "os"
    "strconv"
)

func main() {
    // Bail if there are not enough arguments.
    if len(os.Args) != 2 {
        os.Exit(1)
    }

    // Try to convert the string argument to an integer.
    // Note the second return value which is an error (actually an
    // implementation) of the core error interface.
    n, err := strconv.Atoi(os.Args[1])

    // We can create our own error type; the only requirement is that it
    // fulfills the contract of the built-in error interface.
    // More commonly, we can create our own errors by importing the errors
    // package and using it in the New function:
    //
    // import (
    //     "errors"
    // )
    //
    // func process(count int) error {
    //     if count < 1 {
    //         return errors.New("Invalid count")
    //     }
    //
    //     return nil
    // }

    // If the error is populated then we could not convert the string.
    if err != nil {
        fmt.Println("Not a valid number")
    } else {
        fmt.Println(n)
    }
}
