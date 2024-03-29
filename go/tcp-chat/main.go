package main

import (
	"bufio"
	"log"
	"net"
	"strings"
)

const (
	NETWORK  = "tcp"
	PORT     = ":8080"
	CMD_QUIT = "/quit"
)

func main() {
	chat := NewChat()

	listener, err := net.Listen(NETWORK, PORT)
	if err != nil {
		log.Fatalf("Error starting the server")
	}
	defer listener.Close()

	log.Printf("Server listening on %s", PORT)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Connection failed:", err.Error())
			continue
		}

		log.Println("Connection accepted from", conn.RemoteAddr().String())

		// Prompt the client to enter a username
		conn.Write([]byte("Enter your name: "))
		name, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			name = "Teapot"
		}
		conn.Write([]byte("Type any message to send it, type /quit to finish\n\n"))

		client := NewClient(conn, strings.TrimSpace(name))
		chat.Join(client)
	}
}
