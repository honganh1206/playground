package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", "localhost:9003")
	conn, err := net.DialTCP("tcp", nil, addr)
	// tcpConn := conn.(*net.TCPConn)
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		// Copy output from the server to stdout of the OS
		io.Copy(os.Stdout, conn)
		log.Println("done")
		// signal the main goroutine that the background goroutine is done
		done <- struct{}{}
	}()
	mustCopy(conn, os.Stdin)
	// Close only the write half of the connection
	// if err := tcpConn.CloseWrite(); err != nil {
	// 	panic(err)
	// }
	conn.CloseWrite()
	// Wait for background goroutine to finish,
	// main cannot return when there is no value in done channel
	<-done
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
