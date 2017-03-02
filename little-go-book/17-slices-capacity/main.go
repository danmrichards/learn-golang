package main

import (
	"fmt"
)

// It turns out that append is pretty special. If the underlying array is full,
// it will create a new larger array and copy the values over (this is exactly
// how dynamic arrays work in PHP, Python, Ruby, JavaScript, …).
func main() {
	scores := make([]int, 0, 5)
	c := cap(scores)
	fmt.Println(c)

	for i := 0; i < 25; i++ {
		scores = append(scores, i)

		// If our capacity has changed, Go had to grow our array to accommodate
		// the new data.
		if cap(scores) != c {
			c = cap(scores)
			fmt.Println(c)
		}
	}
}
