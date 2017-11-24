package domain

type Message struct {
	From string
	To   string
	Text string
	Read bool
}

func NewMessage(from, to, text string) *Message {
	message := Message{
		From: from,
		To:   to,
		Text: text,
		Read: false,
	}
	return &message
}
