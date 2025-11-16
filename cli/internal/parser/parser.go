package parser

import (
	"bufio"
	"io"
	"log"
	"net/mail"
	"os"

	"github.com/cravenceiling/indexer/cli/internal/email"
)

type Parser struct{}

var isValidHeader = map[string]bool{
	"Message-Id": true,
	"Date":       true,
	"From":       true,
	"To":         true,
	"Subject":    true,
	"Cc":         true,
	"Bcc":        true,
}

// Parse
func (p Parser) Parse(filePath string) (*email.Email, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	reader := bufio.NewReader(file)

	msg, err := mail.ReadMessage(reader)
	if err != nil {
		return nil, err
	}

	email, err := getEmail(msg)
	if err != nil {
		return nil, err
	}

	return email, nil
}

// getEmail
func getEmail(msg *mail.Message) (*email.Email, error) {
	em := email.Email{}
	var err error

	for k := range msg.Header {
		if isValidHeader[k] {
			em[k] = msg.Header.Get(k)
		}
	}

	content, err := io.ReadAll(msg.Body)
	if err != nil {
		log.Printf("error reading body: %v", err)
	}
	em["Body"] = string(content)

	//buf := &bytes.Buffer{}
	//if _, err = io.Copy(buf, msg.Body); err != nil {
	//	return nil, err
	//}

	//em["Body"] = buf.String()

	return &em, nil
}
