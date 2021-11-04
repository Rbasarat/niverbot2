package mocksource

import (
	"embed"
	"github.com/rbasarat/niverobot/internal/chatbot"
	"log"
	"net"
	"net/http"
	"time"
)

//go:embed *.html
var webFiles embed.FS

type Settings struct {
	Address string
}

type MockSource struct {
	address string
}

func (m *MockSource) Connect() (chan chatbot.Message, error) {
	messages := make(chan chatbot.Message)
	webroot := http.FileServer(http.Dir("internal/mocksource"))
	server := http.Server{Handler: http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodOptions:
			fallthrough
		case http.MethodGet:
			webroot.ServeHTTP(writer, request)
		case http.MethodPost:
			err := request.ParseForm()
			if err != nil {
				http.Error(writer, "must send form posts", 400)
				return
			}
			sender := request.PostFormValue("sender")
			message := request.PostFormValue("message")
			log.Printf("Got mock message: %s - %s", sender, message)
			messages <- chatbot.Message{
				Sender: chatbot.User{
					Username: sender,
				},
				Text: message,
				Date: time.Now(),
			}
		}
	})}

	// Start the server
	address := m.address
	if address == "" {
		address = ":8888"
	}

	listener, err := net.Listen("tcp", address)
	if err != nil {
		return nil, err
	}
	go func() {
		if err := server.Serve(listener); err != http.ErrServerClosed {
			panic(err)
		}
	}()
	log.Printf("started mock service on %s", listener.Addr().String())

	return messages, nil
}

func NewMockSource(settings Settings) chatbot.SourceConfiguration {
	return &MockSource{
		address: settings.Address,
	}
}
