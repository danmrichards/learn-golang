package main

import (
	"fmt"
)

func main() {
	scores := []int{1, 2, 3, 4, 5}
	scores = removeAtIndex(scores, 2)
	fmt.Println(scores)
}

// Removes the item at the given index.
// Achieved by swapping it with the last item and then trimming the slice by 1.
// Only works if we don't care about the order of the slice.
func removeAtIndex(source []int, index int) []int {
	lastIndex := len(source) - 1

	// Swap the last value and the value we want to remove
	source[index], source[lastIndex] = source[lastIndex], source[index]
	return source[:lastIndex]
}
