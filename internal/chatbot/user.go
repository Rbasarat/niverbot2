package chatbot

import "time"

type User struct {
	Id        int
	Username  string
	CreatedAt time.Time
}
