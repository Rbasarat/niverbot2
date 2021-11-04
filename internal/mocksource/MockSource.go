package mocksource

import (
	"embed"
	"github.com/rbasarat/niverobot/internal/chatbot"
	"log"
	"net"
	"net/http"
)

//go:embed *.html
var webFiles embed.FS

type Settings struct {
	Address string
}

func NewMockSource(settings Settings) chatbot.SourceConfiguration {
	webroot := http.FileServer(http.FS(webFiles))
	server := http.Server{Handler: http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodOptions:
			fallthrough
		case http.MethodGet:
			webroot.ServeHTTP(writer, request)
		case http.MethodPost:
			// TODO: CUSTOM CONTROLLER
		}
	})}

	// Start the server
	address := settings.Address
	if address == "" {
		address = ":8888"
	}
	listener, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}
	go func() {
		if err := server.Serve(listener); err != http.ErrServerClosed {
			panic(err)
		}
	}()
	log.Printf("started mock service on %s", listener.Addr().String())

	return nil
}

type MockSource struct {
	address string
	server  http.Server
}
