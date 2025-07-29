package main

import (
	"io"
	"log"
	"net"
	"time"
)

// TIP: We can use killall clock1 to kill clock1

func main() {
	listener, err := net.Listen("tcp", "localhost:9999")
	if err != nil {
		log.Fatal(err)
	}

	for {
		// Block until an incoming connection request is made
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // connection aborted
			continue
		}

		handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}
