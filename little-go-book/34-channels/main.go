package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Channels help make concurrent programming saner by taking shared data out of
// the picture. A channel is a communication pipe between goroutines which is
// used to pass data. In other words, a goroutine that has data can pass it to
// another goroutine via a channel. The result is that, at any point in time,
// only one goroutine has access to the data.

func main() {
	// A channel, like everything else, has a type. This is the type of data
	// that we’ll be passing through our channel. In this case integers.
	c := make(chan int)

	// Start some workers for the channel.
	// We don’t know which worker is going to get what data. What we do know,
	// what Go guarantees, is that the data we send to a channel will only be
	// received by a single receiver.
	for i := 0; i < 5; i++ {
		worker := &Worker{id: i}
		go worker.process(c)
	}

	for {
		// Channels support two operations: receiving and sending. We send to a
		// channel by doing:
		//
		// CHANNEL <- DATA
		//
		// and receive from one by doing:
		//
		// VAR := <-CHANNEL
		//
		// The arrow points in the direction that data flows. In this case we
		// are sending channel c a random integer.
		c <- rand.Int()

		time.Sleep(time.Millisecond * 50)
	}
}

// Worker - A channel worker.
type Worker struct {
	id int
}

// Process a channel item.
// Note the function signature to pass the channel (VAR chan TYPE).
func (w *Worker) process(c chan int) {
	for {
		// Receive the integer from the channel.
		data := <-c

		fmt.Printf("worker %d got %d\n", w.id, data)
	}
}
