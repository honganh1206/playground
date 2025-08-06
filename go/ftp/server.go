package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

type session struct {
	conn     net.Conn
	user     string
	loggedIn bool
	cwd      string
}

// Transfer files from a server to a client
// on a computer network
func main() {
	listener, err := net.Listen("tcp", ":2121") // Non-root
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to accept: %v", err)
			continue
		}
		// Spawn a goroutine for each connection
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	s := &session{conn: conn, cwd: "/"}
	s.writeMsg(220, "Simple Go FTP server ready")

	// Read the commands from io.Reader?
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()
		// Handle commands e.g., "USER anonymous"
		parts := strings.SplitN(line, " ", 2)
		cmd := strings.ToUpper(parts[0])
		arg := ""
		if len(parts) > 1 {
			arg = parts[1]
		}

		if !s.handleCommand(cmd, arg) {
			break
		}

	}
}

func (s *session) handleCommand(cmd, arg string) bool {
	switch cmd {
	case "USER":
		s.user = arg
		s.writeMsg(331, "Username okay, need password")
	case "PASS":
		// For the exercise,
		// in real-world scenario, we need to check credentials
		s.loggedIn = true
		s.writeMsg(230, "User logged in")
	case "PWD":
		s.writeMsg(257, fmt.Sprintf("\"%s\" is current directory", s.cwd))
	case "QUIT":
		s.writeMsg(221, "Goodbye")
		return false
	default:
		s.writeMsg(502, "Command not implemented")
	}
	return true
}

func (s *session) writeMsg(code int, msg string) {
	fmt.Fprintf(s.conn, "%d %s\r\n", code, msg)
}
