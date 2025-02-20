package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring(msg string) <-chan string { // Return receive-only channel
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			// Unpredictable intervals
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

func main() {
	// When main returns, the program exits and takes the boring function down
	// with it
	// We are kind of cheating here: The main function could not see the output
	// from other goroutines
	c := boring("boring!")
	joe := boring("Joe")
	ann := boring("Ann")

	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q\n", <-c) // Receive expression as a value
		fmt.Println(<-joe)
		fmt.Println(<-ann)
	}

	// time.Sleep(2 * time.Second)
	fmt.Println("You're boring; I'm leaving.")
}
