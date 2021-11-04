package main

import (
	"github.com/rbasarat/niverobot/internal/chatbot"
	"github.com/rbasarat/niverobot/internal/mocksource"
	"log"
	"time"
)

type GoDoLater struct {
	Handler chatbot.MessageHandler
}

func (g GoDoLater) Handle(message chatbot.Message) {
	log.Println("I am sleepy, wait a minute")
	go func() {
		time.Sleep(10 * time.Second)
		g.Handler.Handle(message)
	}()

}

func main() {
	chat := chatbot.NewChatbot(chatbot.Settings{
		Configs: []chatbot.SourceConfiguration{
			mocksource.NewMockSource(mocksource.Settings{
				Address: ":5555",
			}),
		},
	})

	logger := chatbot.HandlerFunc(func(message chatbot.Message) {
		log.Printf("Hello ik van %s gekregen: %s....", message.Sender.Username, message.Text)
	})

	butABitLater := GoDoLater{
		Handler: logger,
	}

	chat.Handle(butABitLater)
	chat.Handle(logger)

	chat.Run()
}
