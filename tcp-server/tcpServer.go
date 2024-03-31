package main

import (
	"bufio"
	"log"
	"net"
)

const (
	CONN_TYPE = "tcp"
	CONN_PORT = ":8080"
)

func main() {
	listener, err := net.Listen(CONN_TYPE, CONN_PORT)
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
	defer listener.Close()

	log.Println("Server listening on", CONN_PORT)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error trying to connect:", err.Error())
			break
		}
		log.Println("Accepted connection from", conn.RemoteAddr().String())

		go func() {
			msg, err := bufio.NewReader(conn).ReadString('\n')
			if err != nil {
				log.Println("Error sending message:", err.Error())
			}

			log.Printf("%s -> %s", conn.RemoteAddr().String(), msg)
			conn.Close()
		}()
	}
}
