package main

import (
	"fmt"
)

// In Go, you rarely, if ever, use arrays directly. Instead, you use slices. A
// slice is a lightweight structure that wraps and represents a portion
// of an array.
func main() {
	// Create a slice with values, similar to array syntax.
	// scores := []int{1,4,293,4,9}

	// We use make instead of new because there’s more to creating a slice than
	// just allocating the memory (which is what new does). Specifically, we
	// have to allocate the memory for the underlying array and also initialize
	// the slice. In the below, we initialize a slice with a length of 10
	// and a capacity of 10.
	// scores := make([]int, 10)

	// The length is the size of the slice, the capacity is the size of the
	// underlying array. Using make we can specify the two separately. This
	// creates a slice with a length of 0 but with a capacity of 10.
	scores := make([]int, 0, 10)

	// This will crash the app. The slice has a capacity of 10, but 0 length.
	// We need to expand the array to access the elements.
	// scores[7] = 9033

	// One way to expand is via the append method.
	// scores = append(scores, 5)
	// Prints [5].
	// fmt.Println(scores)

	// But that changes the intent of our original code. Appending to a slice
	// of length 0 will set the first element. For whatever reason, our
	// crashing code wanted to set the element at index 7. To do this, we can
	// re-slice our slice.
	scores = scores[0:8]
	scores[7] = 9033
	fmt.Println(scores)
}
