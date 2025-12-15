---
id: Buffered Channels
aliases: []
tags: []
sr-due: 2025-09-28
sr-interval: 3
sr-ease: 250
---

Tags: #review #golang #programming

Go can have buffered channels (channels holding a certain number of values).

```go
func main() {
	ch := make(chan int, 1)
	ch <- 1
	ch <- 2 // Adding this leads to a deadlock
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
```

## Capacity

> [!IMPORTANT] Buffering removes [synchronization](./Synchronization.md).

We use `cap` for channel's buffer capacity `fmt.Println(cap(ch))` and `len` for the number of elements currently buffered `len(ch)`

In real programs, send and receive operations are usually executed by different goroutines. That's how they communicate. Channels are _deeply connected_ to goroutine scheduling, and without another goroutine receiving from the channel, the sender risks becoming blocked forever.

> If all you need a simple queue, use a slice. Novices are usually tempted to use the buffer channel in a single goroutine as a queue, and that's a mistake.

```go
func mirroredQuery() string {
	responses := make(chan string, 3)
	go func() { responses <- request("asia.gopl.io") }()
	go func() { responses <- request("europe.gopl.io") }()
	go func() { responses <- request("americas.gopl.io") }()
	// Return the quickest response
	// even before the two slower servers have responded
	return <-responses
}

func request(hostname string) (response string) { /* ... */ }
```

## Buffered channels vs. Unbuffered channels

|                    | Buffered                                          | Un-buffered                                  |
| ------------------ | ------------------------------------------------- | -------------------------------------------- |
| Sender -> Receiver | Sender blocks its goroutine if the buffer is full | Sender blocks until receiver ready to read   |
| Receiver -> Sender | Receiver blocks if the buffer is empty            | Receiver blocks until there is value to read |

## Buffering and performance

Channel buffering may impact program performance. Imagine a kitchen with 3 cooks: One bakes, one does the icing, and one inscribes before passing each cake to the next cook in line. _Each cook that has finished a cake must wait for the next cook to become ready to accept it_

Now imagine if there is space for one cake between each cook: A cook may place a finished cake there and immediately starts with the next cake. That is the equivalence of a buffered channel with capacity 1.

Larger space (aka larger buffer) can smooth out bigger transient variations without stalling the assembly line. If one cook decides to catch a break, he can later rush to catch up.

But what if the earlier stages of the assembly line are consistently faster than the later stages? And what if in reverse? The buffer then spends most of it time full, and that buffer provides no benefit here.

And what if the second stage is more elaborate? The cook on the second stage cannot keep up with the supply, so we could hire another cook to help the second one, which is analogous to creating another goroutine communicating over the same channel.
