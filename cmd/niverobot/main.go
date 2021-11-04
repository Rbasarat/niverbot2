package main

import "github.com/rbasarat/niverobot/internal/chatbot"

func main() {
	chat := chatbot.NewChatbot(chatbot.Settings{
		Configs: []chatbot.SourceConfiguration{},
	})

	chat.HandleFunc(func(message chatbot.Message) {

	})

}
