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
			waitForIt := make(chan bool)
			c <- Message{fmt.Sprintf("%s %d", msg, i), waitForIt}
			// Unpredictable intervals
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			<-waitForIt
		}
	}()
	return c
}

// Let whoever is ready to talk talks
func fanIn(inputs ...<-chan Message) <-chan Message {
	c := make(chan Message)

	for _, i := range inputs {
		go func() {
			for msg := range i {
				c <- msg
			}
		}()
	}

	return c
}

func main() {
	// When main returns, the program exits and takes the boring function down
	// with it
	// We are kind of cheating here: The main function could not see the output
	// from other goroutines

	// c := boring("boring!")
	joe := boring("Joe")
	ann := boring("Ann")

	// for i := 0; i < 5; i++ {
	// 	fmt.Printf("You say: %q\n", <-c) // Receive expression as a value
	// 	fmt.Println(<-joe)
	// 	fmt.Println(<-ann)
	// }
	c := fanIn(joe, ann)

	// for i := 0; i < 10; i++ {
	// 	fmt.Println(<-c)
	// }

	for i := 0; i < 5; i++ {
		msg1 := <-c
		fmt.Println(msg1.str)
		msg2 := <-c
		fmt.Println(msg2.str)
		msg1.wait <- true
		msg2.wait <- true

	}

	// time.Sleep(2 * time.Second)
	fmt.Println("You're boring; I'm leaving.")
}
