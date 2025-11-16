package email

import (
	"net/mail"
)

type Email map[string]any

// Email contains all the information of an e-mail.
type CustomEmail struct {
	MessageID string          `json:"messageId"`
	Date      string          `json:"date"`
	From      string          `json:"from"`
	To        []*mail.Address `json:"to"`
	CC        []*mail.Address `json:"cc"`
	BCC       []*mail.Address `json:"bcc"`
	Subject   string          `json:"subject"`
	Body      string          `json:"body"`
}
