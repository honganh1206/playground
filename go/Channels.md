---
id: Channels
aliases: []
tags: []
---

Tags: #review #programming #golang

# Channels in Go

A **typed conduit** through which we send and receive values. It is a way for goroutines to [communicate](./Communication.md)

```go
ch := make(chan int) // Create a channel
ch <- v // Send v to channel ch
v := <-ch // Receive value from ch and assign value to x
```

> [!important] Important Sends and receives are blocked until the other side is ready. This allows
> goroutines to sync without explicit locks or condition variables.

```go
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c) // split the 1st half and sum them
	go sum(s[len(s)/2:], c) // split the second half and sum them
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}

```

We can have _3 types of channels_: Bidirectional, receive-only and send-only

```go
chan time.Time // Bidirectional (can both send and receive)
<-chan time.Time // Receive-only
chan<- time.Time // Send-only
```

[Buffered Channels](./Buffered Channels.md)

## Range and Close

A sender can use the keyword `close` to indicate that no more values will be sent, and receivers can test if the channel has been closed by using a second boolean parameter.

```go

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
	for i := range c { // The loop receives values from the channel repeatedly until it is closed
		fmt.Println(i)
	}
}

```

[[Select statement]]

A channel both [communicates](./Communication.md) and [synchronizes](./Synchronization.md)
