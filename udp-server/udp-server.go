package main

import (
	"fmt"
	"log"
	"net"
)

const (
	CONN_TYPE = "udp"
	CONN_PORT = ":8080"
)

func main() {
	addr, err := net.ResolveUDPAddr(CONN_TYPE, CONN_PORT)
	if err != nil {
		log.Fatalf("Error resolving address: %s", err.Error())
	}

	conn, err := net.ListenUDP(CONN_TYPE, addr)
	if err != nil {
		log.Fatalf("Error listening: %s", err.Error())
	}
	defer conn.Close()

	log.Println("Server listening on", CONN_PORT)

	buffer := make([]byte, 1024)
	for {
		n, addr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			log.Println("Error trying to read:", err.Error())
			continue
		}

		message := string(buffer[:n])
		log.Printf("%s -> %s", addr.String(), message)

		go func() {
			response := fmt.Sprintf("Server received: %s", message)
			_, err := conn.WriteToUDP([]byte(response), addr)
			if err != nil {
				log.Println("Error writing response:", err.Error())
			}
		}()
	}
}
