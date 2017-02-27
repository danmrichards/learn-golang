package main

import (
    "fmt"
    "math/rand"
    "sort"
)

func main() {
    // Create slice with 100 length and capacity.
    scores := make([]int, 100)

    // Populate with random values.
    for i := 0; i < 100; i++ {
        scores[i] = int(rand.Int31n(1000))
    }

    // Sort ASC.
    sort.Ints(scores)

    // Create slice with 5 capacity.
    worst := make([]int, 5)

    // Get the lowest 5 scores.
    // copy(worst, scores[:5])

    // Get the lowest 5 scores but add them to indexes 2 to 4 of the slice.
    // copy(worst[2:4], scores[:5])

    // Get the lowest 3 scores.
    // copy(worst, scores[:3])

    fmt.Println(worst)
}
