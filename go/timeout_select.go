package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Message struct {
	str  string
	wait chan bool // Dedicated wait channel per message
}

func boring(msg string) <-chan Message { // Return receive-only channel
	c := make(chan Message)
	go func() {
		for i := 0; ; i++ {
			waitForIt := make(chan bool) // Sync mechanism between sender and receiver
			c <- Message{fmt.Sprintf("%s %d", msg, i), waitForIt}
			// Unpredictable intervals
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			<-waitForIt // The sender then waits on this
		}
	}()
	return c
}

func main() {
	c := boring("Joe")
	timeout := time.After(5 * time.Second)
	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-timeout:
			// Timeout the entire conversation
			fmt.Println("You talk too much.")
			return
		}
	}
}
