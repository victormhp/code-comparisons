package main

import (
	"fmt"
	"log"
	"strings"
)

type Chat struct {
	clients  []*Client
	join     chan *Client
	leave    chan *Client
	incoming chan *Message
	messages []string
}

func NewChat() *Chat {
	chat := &Chat{
		clients:  make([]*Client, 0),
		join:     make(chan *Client),
		leave:    make(chan *Client),
		incoming: make(chan *Message),
	}

	chat.Listen()
	return chat
}

func (chat *Chat) Listen() {
	go func() {
		for {
			select {
			case message := <-chat.incoming:
				chat.Parse(message)
			case client := <-chat.join:
				chat.Join(client)
			case client := <-chat.leave:
				chat.Leave(client)
			}
		}
	}()
}

func (chat *Chat) Join(client *Client) {
	chat.clients = append(chat.clients, client)
	chat.Broadcast(fmt.Sprintf("%s joined the chat!\n", client.name))

	go func() {
		for message := range client.incoming {
			chat.incoming <- message
		}
		chat.leave <- client
	}()
}

func (chat *Chat) Leave(client *Client) {
	for i, c := range chat.clients {
		if c == client {
			chat.clients = append(chat.clients[:i], chat.clients[i+1:]...)
			break
		}
	}
	close(client.outgoing)
	log.Printf("%s has disconnected\n", client.conn.RemoteAddr().String())
	chat.Broadcast(fmt.Sprintf("%s leaved the chat!\n", client.name))
}

func (chat *Chat) Parse(message *Message) {
	switch {
	default:
		message.client.ClearLine()
		chat.Broadcast(message.String())
	case strings.HasPrefix(message.text, CMD_QUIT):
		message.client.Quit()
	}
}

func (chat *Chat) Broadcast(message string) {
	for _, client := range chat.clients {
		client.outgoing <- message
	}
}
