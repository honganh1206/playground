package main

import (
	"fmt"
	"sync"
	"time"
)

// Combine multiple channels into one
func fanIn(cs ...<-chan string) <-chan string {
	merged := make(chan string)
	// To manage channels properly
	var wg sync.WaitGroup
	wg.Add(len(cs))

	for _, c := range cs {
		go func() {
			for msg := range c {
				merged <- msg
			}
			wg.Done()
		}()
	}

	go func() {
		// This will cause blocking if it stays in the same goroutine as the one calling fanIn()
		wg.Wait()
		close(merged) // Can only return merged if both c1 and c2 are drained
	}()

	return merged // Can only be executed after waiting
}

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	merged := fanIn(c1, c2)

	go func() {
		for i := 0; i < 5; i++ {
			c1 <- fmt.Sprintf("Message from c1: %d", i)
			time.Sleep(100 * time.Millisecond)
		}
		close(c1)
	}()

	go func() {
		for i := 0; i < 5; i++ {
			c2 <- fmt.Sprintf("Message from c2: %d", i)
			time.Sleep(150 * time.Millisecond)
		}
		close(c2)
	}()

	for msg := range merged {
		fmt.Println(msg)
	}

	fmt.Println("All messages received")
}
