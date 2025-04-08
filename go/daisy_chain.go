package main

import "fmt"

// Receive from right, add 1, send to left
func f(left, right chan int) {
	left <- 1 + <-right
}

// Each channel receives input from its predecessor channel
// It then passes output to its successive channel
// Like a line of people passing a message
func main() {
	const n = 10000
	leftmost := make(chan int)
	right := leftmost
	left := leftmost

	// leftmost = chan0 ← chan1 ← chan2 ← chan3 ← ... ← chanN
	for i := 0; i < n; i++ {
		right = make(chan int)
		go f(left, right)
		left = right
	}

	// 	Initial value (1) → chanN → chanN-1 → ... → chan1 → leftmost
	// (Each step adds 1 to the value)
	go func(c chan int) { c <- 1 }(right)
	fmt.Println(<-leftmost)
}
