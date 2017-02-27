package main

import (
    "fmt"
    "strings"
)

func main() {
    haystack := "The spice must flow"
    spaceIndex := strings.Index(haystack[:len(haystack)], " ")
    spaceIndexAfterFive := strings.Index(haystack[5:], " ")

    fmt.Printf("The first space was at index %d\n", spaceIndex)
    fmt.Printf("The first space after 5 characters was at index %d\n", spaceIndexAfterFive)
}
