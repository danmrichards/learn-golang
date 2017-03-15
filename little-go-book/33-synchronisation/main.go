package main

import (
	"fmt"
	"sync"
	"time"
)

// The only concurrent thing you can safely do to a variable is to read from it.
// You can have as many readers as you want, but writes need to be synchronized.
// There are various ways to do this, including using some truly atomic
// operations that rely on special CPU instructions. However, the most common
// approach is to use a mutex.

// While it might be tempting to use coarse locks (locks that cover a large
// amount of code), that undermines the very reason we’re doing concurrent
// programming in the first place. We generally want fine locks; else, we end up
// with a ten-lane highway that suddenly turns into a one-lane road.

var (
	counter = 0
	lock    sync.Mutex
)

func main() {
	for i := 0; i < 20; i++ {
		go incr()
	}
	time.Sleep(time.Millisecond * 10)
}

// A mutex serializes access to the code under lock. The reason we simply define
// our lock as lock sync.Mutex is because the default value of a sync.Mutex
// is unlocked.
func incr() {
	lock.Lock()
	defer lock.Unlock()

	counter++
	fmt.Println(counter)
}
