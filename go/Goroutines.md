---
id: Goroutines
aliases: []
tags: []
---

# Goroutines

A lightweight thread managed by the Go runtime. We can think of it as the async-await mechanism in
C#: A little helper we can call to do something for us in the background while we are busy with
other stuff.

```go
go f(x, y, z)
// Start a new gorountine
// We evaluate f, x, y , z during the CURRENT gorountine, while the execution of f happens in the NEW gorountine
f(x, y, z)

// Example code
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

func main() {
	go say("world") // Start a helper (goroutine) to say "world"
	say("hello")    // Main program says "hello" directly
}
```

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
> Goroutines are NOT threads. There might only be one thread in a program, and that thread has thousands of goroutines.

Goroutines are **multiplexed dynamically** onto threads, and thinking of goroutines as cheap threads will not take you far.
