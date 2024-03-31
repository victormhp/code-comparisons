package main

import (
	"bufio"
	"log"
	"net"
)

type Client struct {
	conn     net.Conn
	name     string
	incoming chan *Message
	outgoing chan string
	reader   *bufio.Reader
	writer   *bufio.Writer
}

func NewClient(conn net.Conn, name string) *Client {
	client := &Client{
		conn:     conn,
		name:     name,
		incoming: make(chan *Message),
		outgoing: make(chan string),
		reader:   bufio.NewReader(conn),
		writer:   bufio.NewWriter(conn),
	}

	client.Listen()
	return client
}

func (client *Client) Listen() {
	go client.Read()
	go client.Write()
}

func (client *Client) Read() {
	for {
		str, err := client.reader.ReadString('\n')
		if err != nil {
			break
		}

		if len(str) > 1 {
			message := NewMessage(client, str)
			client.incoming <- message
		}
	}
	close(client.incoming)
}

func (client *Client) Write() {
	for str := range client.outgoing {
		_, err := client.writer.WriteString(str)
		if err != nil {
			log.Println("Error:", err.Error())
			break
		}

		err = client.writer.Flush()
		if err != nil {
			log.Println("Error:", err.Error())
			break
		}
	}
}

func (client *Client) ClearLine() {
	// ANSI escape to move up a line: \033[A
	// ANSI escape to clear the line: \033[2K
	_, err := client.conn.Write([]byte("\033[A\033[2K"))
	if err != nil {
		log.Println("Error clearing prompt:", err.Error())
	}
}

func (client *Client) Quit() {
	client.conn.Close()
}
