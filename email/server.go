package email

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func handleConnection(conn net.Conn) {
	fmt.Println("New connection from", conn.RemoteAddr())
	defer conn.Close()

	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	// Send greeting
	writer.WriteString("220 smtp.example.com Simple Mail Transfer Service Ready\r\n")
	writer.Flush()

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Failed to read client message:", err)
			return
		}
		line = strings.TrimSpace(line)

		switch {
		case strings.HasPrefix(line, "HELO") || strings.HasPrefix(line, "EHLO"):
			writer.WriteString("250 smtp.example.com at your service\r\n")
		case strings.HasPrefix(line, "MAIL FROM:"):
			writer.WriteString("250 2.1.0 Sender OK\r\n")
		case strings.HasPrefix(line, "RCPT TO:"):
			writer.WriteString("250 2.1.5 Recipient OK\r\n")
		case strings.HasPrefix(line, "DATA"):
			writer.WriteString("354 Start mail input; end with <CRLF>.<CRLF>\r\n")
			writer.Flush()

			bodyLines, err := readEmailBody(reader)
			if err != nil {
				fmt.Println("Failed to read email body:", err)
				return
			}
			body := strings.Join(bodyLines, "\n")

			// Parse the email and convert it to JSON
			parsedEmail, err := parseEmail(body)
			if err != nil {
				fmt.Println("Failed to parse email:", err)
				return
			}
			fmt.Printf("Email From %s\n", parsedEmail.From)
			fmt.Printf("Email To %s\n", parsedEmail.To)
			fmt.Printf("Email Subject %s\n", parsedEmail.Subject)

			if err != nil {
				fmt.Println("Failed to print email as formatted JSON:", err)
				return
			}
			writer.WriteString("250 2.0.0 OK\r\n")
		case strings.HasPrefix(line, "QUIT"):
			writer.WriteString("221 2.0.0 Bye\r\n")
			writer.Flush()
			return
		default:
			writer.WriteString("500 5.5.1 Command unrecognized\r\n")
		}
		writer.Flush()
	}
}

func StartServer() {
	ln, err := net.Listen("tcp", ":25")
	if err != nil {
		fmt.Println("Failed to start server:", err)
		return
	}
	defer ln.Close()
	fmt.Println("SMTP Server is online!")

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Failed to accept connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}

func ShutdownServer() {
	fmt.Println("SMTP Server is offline!")
}
