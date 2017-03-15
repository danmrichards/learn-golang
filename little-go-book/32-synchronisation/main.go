package main

import (
	"fmt"
	"time"
)

// If you think the output is 1, 2, ... 20 you’re both right and wrong. It’s
// true that if you run the above code, you’ll sometimes get that output.
// However, the reality is that the behavior is undefined. Why? Because we
// potentially have multiple (two in this case) goroutines writing to the same
// variable, counter, at the same time. Or, just as bad, one goroutine would be
// reading counter while another writes to it.

// The only concurrent thing you can safely do to a variable is to read from it.
// You can have as many readers as you want, but writes need to be synchronized.
// There are various ways to do this, including using some truly atomic
// operations that rely on special CPU instructions.

var counter = 0

func main() {
	for i := 0; i < 20; i++ {
		go incr()
	}
	time.Sleep(time.Millisecond * 10)
}

func incr() {
	counter++
	fmt.Println(counter)
}
