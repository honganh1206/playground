---
id: Goroutine leaks
aliases:
  - Goroutine leaks
tags: []
---

# Goroutine leaks

```go
func mirroredQuery() string {
	responses := make(chan string)
	go func() { responses <- request("asia.gopl.io") }()
	go func() { responses <- request("europe.gopl.io") }()
	go func() { responses <- request("americas.gopl.io") }()
	// Return the quickest response
	// even before the two slower servers have responded
	return <-responses
}

func request(hostname string) (response string) { /* ... */ }
```

In the above example, we use an unbuffered channel, and the main goroutine is the receiver.

When the quickest response returns, the two slower goroutines will be **stuck** trying to send their responses as `mirroedQuery` function has exited i.e., no receiver (Sender blocks until receiver is ready to receive, per [[Buffered Channels]]).

This is called a **goroutine leak**. The two slower goroutines are now stuck in memory and can consume resources indefinitely.

Unlike garbage variables, leaked goroutines are not automatically collected, so we need to ensure _goroutines can terminate themselves when no longer needed_.
