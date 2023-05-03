package email

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/mail"
	"strings"
)

type ParsedEmail struct {
	From        string            `json:"from"`
	To          string         `json:"to"`
	Cc          []string          `json:"cc"`
	Bcc         []string          `json:"bcc"`
	Subject     string            `json:"subject"`
	Body        string            `json:"body"`
	ContentType string            `json:"content_type"`
	Headers     map[string]string `json:"headers"`
}

func parseEmail(rawEmail string) (*ParsedEmail, error) {
	msg, err := mail.ReadMessage(strings.NewReader(rawEmail))
	if err != nil {
		return nil, fmt.Errorf("failed to parse email: %w", err)
	}

	from, _ := msg.Header.AddressList("From")
	to, _ := msg.Header.AddressList("To")
	cc, _ := msg.Header.AddressList("Cc")
	bcc, _ := msg.Header.AddressList("Bcc")

	headers := make(map[string]string)
	for key := range msg.Header {
		headers[key] = msg.Header.Get(key)
	}

	body := ""
	contentType := msg.Header.Get("Content-Type")

	if contentType != "multipart/mixed" {
		bodyBytes, err := ioutil.ReadAll(msg.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to read email body: %w", err)
		}
		body = string(bodyBytes)
	} else {
		// You can handle multipart messages and attachments here.
		// For the sake of simplicity, this example does not cover this part.
	}

	parsedEmail := &ParsedEmail{
		From:        from[0].Address,
		To:          to[0].Address,
		Cc:          addressesToStrings(cc),
		Bcc:         addressesToStrings(bcc),
		Subject:     msg.Header.Get("Subject"),
		Body:        body,
		ContentType: contentType,
		Headers:     headers,
	}

	return parsedEmail, nil
}


func readEmailBody(reader *bufio.Reader) ([]string, error) {
	var bodyLines []string
	for {
		bodyLine, err := reader.ReadString('\n')
		if err != nil {
			return nil, fmt.Errorf("failed to read email body: %w", err)
		}
		bodyLine = strings.TrimRight(bodyLine, "\r\n")

		if bodyLine == "." {
			break
		}

		bodyLines = append(bodyLines, bodyLine)
	}
	return bodyLines, nil
}


func addressesToStrings(addresses []*mail.Address) []string {
	stringsList := make([]string, 0, len(addresses))
	for _, addr := range addresses {
		stringsList = append(stringsList, addr.String())
	}
	return stringsList
}

func emailToJSON(parsedEmail *ParsedEmail) (string, error) {
	jsonBytes, err := json.Marshal(parsedEmail)
	if err != nil {
		return "", fmt.Errorf("failed to convert email to JSON: %w", err)
	}
	return string(jsonBytes), nil
}
