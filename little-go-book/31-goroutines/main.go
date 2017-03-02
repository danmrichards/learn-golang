package main

import (
	"fmt"
	"time"
)

// A goroutine is similar to a thread, but it is scheduled by Go, not the OS.
// Code that runs in a goroutine can run concurrently with other code.

// Goroutines are easy to create and have little overhead. Multiple goroutines
// will end up running on the same underlying OS thread. This is often called an
// M:N threading model because we have M application threads (goroutines)
// running on N OS threads. The result is that a goroutine has a fraction of
// overhead (a few KB) than OS threads. On modern hardware, it’s possible to
// have millions of goroutines. Furthermore, the complexity of mapping and
// scheduling is hidden. We just say this code should run concurrently and let
// Go worry about making it happen.

func main() {
	fmt.Println("Start")

	// We simply use the go keyword followed by the function we want to execute.
	// If we just want to run a bit of code, such as the below, we could use an
	// anonymous function. Do note that anonymous functions aren’t only used
	// with goroutines, however. Example:
	//
	// go func() {
	//     fmt.Println("processing")
	// }()
	go process()

	// This is bad. Don't do this!
	// We're doing this because the main process exits before the goroutine gets
	// a chance to execute (the process doesn’t wait until all goroutines are
	// finished before exiting).
	time.Sleep(time.Millisecond * 10)
	fmt.Println("Done")
}

func process() {
	fmt.Println("Processing")
}
