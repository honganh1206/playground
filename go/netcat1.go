package main

import (
	"io"
	"log"
	"net"
	"os"
)

// Manipulate network connections
// Read data from a connection and write to stdout
// until EOF condition or an error occurs

func main() {
	conn, err := net.Dial("tcp", "localhost:9999")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
