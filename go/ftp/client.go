package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:2121")
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	reader := bufio.NewReader(conn)

	greet, _ := reader.ReadString('\n')
	fmt.Print(greet)

	// Read user input
	stdinReader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("ftp>")
		cmdLine, _ := stdinReader.ReadString('\n')
		cmdLine = strings.TrimSpace(cmdLine)

		if cmdLine == "" {
			continue
		}

		fmt.Fprintf(conn, "%s\r\n", cmdLine)

		// Read server response line by line
		resp, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Connection closed.")
			return
		}
		fmt.Print(resp)

		if strings.HasPrefix(strings.ToUpper(cmdLine), "QUIT") {
			break
		}
	}
}
