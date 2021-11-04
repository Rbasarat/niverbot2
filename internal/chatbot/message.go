package chatbot

import "time"


type Message struct {
	text string
	sender User
	date time.Time
}