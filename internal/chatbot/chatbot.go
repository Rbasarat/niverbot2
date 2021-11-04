package chatbot

type Chatbot struct{}

type Settings struct {
	Configs []SourceConfiguration
}

type SourceConfiguration interface {
}

func NewChatbot(settings Settings) *Chatbot {
	return nil
}

type MessageHandler interface {
	Handle(message Message)
}

func (c *Chatbot) Handle(handler MessageHandler) {

}

type HandlerFunc func(message Message)

func (t HandlerFunc) Handle(message Message) {
	t(message)
}

func (c *Chatbot) HandleFunc(handler func(message Message)) {
	c.Handle(HandlerFunc(handler))
}
