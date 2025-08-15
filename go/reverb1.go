package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

// An echo server that uses multiple goroutines per connection

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
	time.Sleep(delay)
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		// Make a real echo
		// consisting of the composition of three independent shouts
		// i.e., making the echoes concurrent and overlap in time
		go echo(c, input.Text(), 1*time.Second)
	}
	// Ignore potential errors from input.Err()
	c.Close()
}

func main() {
	listener, err := net.Listen("tcp", "localhost:9003")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		handleConn(conn) // handle one connection at a time
	}
}
