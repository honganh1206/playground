---
id: Synchronization
aliases:
  - Synchronization
tags: []
---

# Synchronization

It is like how [[Channels]] works: When the main function executes `<-c`, it will wait for a value
to be sent

When the invoked function executes `c <- value`, it will wait for a receiver to be ready

A sender and receiver must **both be ready** to play their part. Otherwise we wait until they are

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
