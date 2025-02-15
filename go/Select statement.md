---
id: Select statement
aliases: []
tags: []
---

Tags: #review #go #programming

Allow a goroutine to _wait on multiple communication operations_

`select` blocks until one of the cases can run, then it execute that case. If multiple cases are ready, it will choose a random case.

```go
func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit: // Will keep running until reaching this case
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int) // This channel helps synchronizing the two gorountines
	quit := make(chan int)
	go func() { // 2nd gorountine - This runs concurrently with the 1st goroutine
		for i := 0; i < 10; i++ {
			fmt.Println(<-c) // Receive value from channel c
		}
		quit <- 0
	}()
	fibonacci(c, quit) // 1st goroutine - This starts first
}
```

The `default` case will be run if no other case is ready

```go
select {
case i := <-c:
    // use i
default:
    // receiving from c would block
}

```
