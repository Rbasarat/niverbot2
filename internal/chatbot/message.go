package chatbot

import "time"

type Message struct {
	Text   string
	Sender User
	Date   time.Time
}
