package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			// Send the current Fibonacci number through channel 'c'
			// This case is selected when the receiver (anonymous goroutine) is ready to receive
			x, y = y, x+y
		case <-quit:
			// Receive signal from quit channel and terminate the function
			// This happens after all Fibonacci numbers have been generated
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int) // This channel helps synchronizing the two goroutines
	quit := make(chan int)
	// Anonymous goroutine
	go func() { // 2nd goroutine - This runs concurrently with the 1st goroutine
		for i := 0; i < 10; i++ {
			fmt.Println(<-c) // Receive value from channel c
		}
		// Signal to the fibonacci function that we're done
		quit <- 0
	}()
	// Start generating Fibonacci numbers in the main goroutine
	// This will block until the anonymous goroutine completes
	fibonacci(c, quit)
}
