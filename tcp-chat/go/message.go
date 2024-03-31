package main

import "fmt"

type Message struct {
	client *Client
	text   string
}

func NewMessage(client *Client, text string) *Message {
	return &Message{
		client: client,
		text:   text,
	}
}

func (message *Message) String() string {
	return fmt.Sprintf("[%s] -> %s", message.client.name, message.text)
}
