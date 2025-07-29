---
id: Concurrency in Go
aliases: []
tags: []
---

[[Goroutines]]

[[Channels]]

[[Mutexes]]

[[Concurrency Pattern]]

[Slide by Rob Pike](https://go.dev/talks/2012/concurrency.slide)

## What is concurrency?

Concurrency is the **composition** of _independently executing computations_

Concurrency mirrors a complex world of interacting, independently behaving pieces

[[Concurrency is NOT Parallelism]], but it **enables** parallelism. If we have only one processor,
we can enable concurrency but not parallelism

A well-written concurrent program can run _efficiently_ in parallel on a multiprocessor

Some languages sharing the concurrency feature are Occam, Erlang, Newsqueak, etc.,

## Concurrency in Go

Go enables **two styles of concurrency programming**: Communicating sequential processes or CSP (goroutines and channels) and **shared memory multithreading**, which is more traditional

> Go's support for concurrency is one of its great strengths, but concurrent programs are inherently harder to understand compared to sequential ones.

In Go, **channels are first-class values**

`go` is analogous to the `&` on the end of a shell command

```bash
# Running a command with &
sleep 5 &  # This runs in the background
echo "This prints immediately"  # This runs without waiting for sleep to finish
```

In Go:

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    // Using goroutine (similar to & in shell)
    go func() {
        time.Sleep(5 * time.Second)
        fmt.Println("5 seconds passed")
    }()

    fmt.Println("This prints immediately")
    time.Sleep(6 * time.Second) // Keep main program running
}
```

[[Communication]]

[[Synchronization]]

[[Multiplexing]]

[[Writing system software with Go]]
