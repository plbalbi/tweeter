package domain

type Message struct {
	From string
	Text string
	Read bool
}

func NewMessage(from, text string) *Message {
	message := Message{
		From: from,
		Text: text,
		Read: false,
	}
	return &message
}
