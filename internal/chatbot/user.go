package chatbot

import "time"

type User struct {
	id        int
	username  string
	createdAt time.Time
}
