package models

// An Email contains all the information of an e-mail.
type Email struct {
	MessageID string `json:"messageId"`
	Date      string `json:"date"`
	From      string `json:"from"`
	To        string `json:"to"`
	CC        string `json:"cc"`
	BCC       string `json:"bcc"`
	Subject   string `json:"subject"`
	Body      string `json:"body"`
}
