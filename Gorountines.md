Tags: #review 

A lightweight thread managed by the Go runtime.  We can think of it as the async-await mechanism in C#: A little helper we can call to do something for us in the background while we are busy with other stuff.


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

As gorountines *run in the same address space*, the access to shared memory must be synchronized. 