package main

type Request struct {
	fn func() int // Operation to perform
	c  chan int   // Channel to return the result
}

func main() {
}
