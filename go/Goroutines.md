---
id: Goroutines
aliases: []
tags: []
---

# Goroutines

## What is a goroutine?

A lightweight thread managed by the Go runtime. We can think of it as the async-await mechanism in
C#: A little helper we can call to do something for us in the background while we are busy with other stuff.

A sequential program may call one function and then call the other, but in a concurrent Go program, _calls to both functions can be active at the same time_

```go
package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond) // Waits for 100ms before continuing
		fmt.Println(s)                    // Prints the string s
	}
}

// main is its own goroutine
// When main returns, all goroutines are abruptly terminated
func main() {
	go say("world") // Start a helper (goroutine) to say "world"
	say("hello")    // Main program says "hello" directly
}
```

> Other than by returning from `main` or exiting the program, no goroutine can stop another one from executing.

## How does a goroutine work?

Goroutines have their own call stack, which grows and shrinks as required

As gorountines _run in the same address space_, the access to shared memory must be synchronized.

An example of a background goroutine check:

```go
// Background goroutine to delete not-seen-recently clients
go func() {
	for {
		time.Sleep(time.Minute)
		// Lock to prevent check limiter check while the cleaning is taking place
		mu.Lock()
		for ip, client := range clients {
			// Remove clients if they are not seen more than 3 mins
			if time.Since(client.lastSeen) > 3*time.Minute {
				delete(clients, ip)
			}
		}
		mu.Unlock()
	}
}() // Immediate function call operator

```

Note that _background goroutine forms a closure over variables outside its scope_. Any changes we
make to the variables will be reflected in the rest of the codebase

```go
func main() {
	exampleString := "foo"

	go func() {
		exampleString += "bar"
		fmt.Println(exampleString) // Prints "foobar"
	}()

	time.Sleep(200 * time.Millisecond)

	fmt.Println(exampleString) // Also prints "foobar"
}
```

> [!WARNING]
> Goroutines are NOT threads. There might only be one thread in a program, and that thread has thousands of goroutines. The differences between threads and goroutines are essentially quantitative.

Networking is a natural domain that benefits a lot from concurrency, since a server typically handles multiple connections from their client. See more at [[The net package]]

## Goroutines and threads

Goroutines are **multiplexed dynamically** onto threads, and thinking of goroutines as cheap threads will not take you far.

## Example with `clock1` and `clock2`

We have `clock1` as a sequential clock, while `clock2` is concurrent

We also have `netcat1` that reads data from a connection and writes it to the standard output.

We run two `netcat1` instances at the same time on different terminals. When we terminate the 1st instance, _the 2nd instance will run_ since the server is sequential - it deals with one client at a time
