package types

import "time"

type Message struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Message string `json:"message"`
	Date    string `json:"date"` // RFC3339 format
}

type DBMessage struct {
	Sender    string
	Receiver  string
	Message   string
	Timestamp time.Time
}
