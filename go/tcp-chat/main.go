package main

import (
	"bufio"
	"fmt"
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

		go func() {
			// Prompt the client to enter a username
			conn.Write([]byte("Enter your name: "))
			name, err := bufio.NewReader(conn).ReadString('\n')
			if err != nil {
				name = "Teapot"
			}

			name = strings.TrimSpace(name)
			conn.Write([]byte(fmt.Sprintf("\nWelcome %s!\n", name)))
			conn.Write(
				[]byte(fmt.Sprintf("Type any message to send it, type %s to finish\n\n", CMD_QUIT)),
			)

			client := NewClient(conn, name)
			chat.Join(client)
		}()
	}
}
