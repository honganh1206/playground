package main

import (
	"fmt"
	"sync"
)

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()

	return out
}

// Same type for its inbound and outbound channels
func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

// Convert a list of channels to a single channel

func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	// Start an output goroutine for each input channel
	// Copy values from c to out until c is closed
	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done() // Signal goroutine completion
	}

	// Track active goroutines
	wg.Add(len(cs))
	// Start a goroutine for each inbound channel
	// that copies the values to the sole outbound channel
	for _, c := range cs {
		go output(c)
	}

	// Close out channel once all the output goroutines are done
	// This must start after the wg.Add call
	go func() {
		wg.Wait() // Wait for all goroutines to finish
		close(out)
	}()

	return out
}

func main() {
	// // Set up the pipeline and consume the output
	// for n := range sq(sq(gen(2, 3))) {
	// 	fmt.Println(n)
	// }
	in := gen(2, 3)

	// Distribute the sq work across two goroutines
	c1 := sq(in)
	c2 := sq(in)

	// Consume the merged output from c1 and c2
	for n := range merge(c1, c2) {
		fmt.Println(n)
	}
}
