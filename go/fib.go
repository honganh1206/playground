package main

import "fmt"

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x // Send x to channel c
		x, y = y, x+y
	}
	close(c) // Only a sender should close a channel
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c) // cap stands for capacity?
	for i := range c {      // The loop receives values from the channel repeatedly until it is closed
		fmt.Println(i)
	}
}
