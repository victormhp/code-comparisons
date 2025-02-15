package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

const (
	CONN_TYPE = "tcp"
	CONN_PORT = ":8080"
)

func main() {
	ln, err := net.Listen(CONN_TYPE, CONN_PORT)
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
	defer ln.Close()

	log.Println("Server listening on", CONN_PORT)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("Error trying to connect:", err.Error())
			continue
		}
		log.Println("Accepted connection from", conn.RemoteAddr().String())

		go func() {
			defer conn.Close()

			for {
				msg, err := bufio.NewReader(conn).ReadString('\n')
				if err != nil {
					log.Println("Error receiving message:", err.Error())
					return
				}
				log.Printf("%s -> %s", conn.RemoteAddr().String(), msg)

				response := fmt.Sprintf("Server received: %s", msg)
				_, err = conn.Write([]byte(response))
				if err != nil {
					log.Println("Error writing response:", err.Error())
					return
				}
			}
		}()
	}
}
