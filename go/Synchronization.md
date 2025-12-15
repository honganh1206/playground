---
id: Synchronization
aliases:
  - Synchronization
tags: []
---

# Synchronization

It is like how [[Channels]] works: When the main function executes `<-c`, it will wait for a value to be sent

When the invoked function executes `c <- value`, it will wait for a receiver to be ready

A sender and receiver must **both be ready** to play their part. Otherwise, we wait until they are.

[[Unbuffered channel]] are synchronous: When a value is sent on an unbuffered channel, the receipt of the value _happens before the reawakening_ of the sending goroutine (More in the [[Go memory model]]). This means two things:

1. All writes made before the send in the sender goroutine become visible to the receiver
2. The receiver gets the value first, then the sender is allowed to resume

```go
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
```
