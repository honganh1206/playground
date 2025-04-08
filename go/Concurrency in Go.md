---
id: Concurrency in Go
aliases: []
tags: []
---

[[Goroutines]]

[[Channels]]

[[Mutexes]]

[[Concurrency Pattern]]

# From the [slide](https://go.dev/talks/2012/concurrency.slide)

Concurrency is the **composition** of _independently executing computations_

Concurrency mirrors a complex world of interacting, independently behaving pieces

[[Concurrency is NOT Parallelism]], but it **enables** parallelism. If we have only one processor,
we can enable concurrency but not parallelism

A well-written concurrent program can run _efficiently_ in parallel on a multiprocessor

Some languages sharing the concurrency feature are Occam, Erlang, Newsqueak, etc.,

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
