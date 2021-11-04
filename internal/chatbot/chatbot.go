package chatbot

type Chatbot struct {
	messages chan Message
	handlers []MessageHandler
}

type Settings struct {
	Configs []SourceConfiguration
}

type SourceConfiguration interface {
	Connect() (chan Message, error)
}

func NewChatbot(settings Settings) *Chatbot {

	messages := make(chan Message)

	// Merge all sources
	sources := make([]chan Message, 0)
	for _, config := range settings.Configs {
		sourceMessages, err := config.Connect()
		if err != nil {
			panic(err)
		}
		sources = append(sources, sourceMessages)
	}
	for _, source := range sources {
		go func() {
			for message := range source {
				messages <- message
			}
		}()
	}

	return &Chatbot{
		handlers: make([]MessageHandler, 0),
		messages: messages,
	}
}

type MessageHandler interface {
	Handle(message Message)
}

func (c *Chatbot) Handle(handler MessageHandler) {
	c.handlers = append(c.handlers, handler)
}

type HandlerFunc func(message Message)

func (h HandlerFunc) Handle(message Message) {
	h(message)
}

func (c *Chatbot) HandleFunc(handler func(message Message)) {
	c.Handle(HandlerFunc(handler))
}

func (c *Chatbot) Run() {
	for message := range c.messages {
		c.fireMessage(message)
	}
}

func (c *Chatbot) fireMessage(message Message) {
	for _, handler := range c.handlers {
		go handler.Handle(message)
	}
}
