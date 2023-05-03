package email

import (
	"fmt"
	"net"
	"strings"
)

func handleConnection(conn net.Conn) {
	fmt.Println("New connection from", conn.RemoteAddr())
	defer conn.Close()

	// Read the client's message
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Failed to read client message:", err)
		return
	}

	// Extract the sender and recipients from the message
	message := string(buf[:n])
	lines := strings.Split(message, "\r\n")
	if len(lines) < 2 {
		fmt.Println("Invalid message from client:", message)
		return
	}
	from := strings.TrimPrefix(lines[0], "MAIL FROM:<")
	from = strings.TrimSuffix(from, ">")
	recipients := []string{}
	for _, line := range lines[1:] {
		if strings.HasPrefix(line, "RCPT TO:<") {
			rcpt := strings.TrimPrefix(line, "RCPT TO:<")
			rcpt = strings.TrimSuffix(rcpt, ">")
			recipients = append(recipients, rcpt)
		}
	}

	// Print the sender and recipients
	fmt.Println("Received message from", from)
	fmt.Println("Recipients:", recipients)

	// Send a response to the client
	conn.Write([]byte("250 OK\r\n"))
}

func StartServer() {
	// Listen on port 25
	ln, err := net.Listen("tcp", ":25")
	if err != nil {
		fmt.Println("Failed to start server:", err)
		return
	}
	defer ln.Close()

	// Accept connections
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Failed to accept connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}