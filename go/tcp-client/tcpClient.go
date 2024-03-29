package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const (
	CONN_TYPE = "tcp"
	CONN_PORT = ":8080"
)

func main() {
	conn, err := net.Dial(CONN_TYPE, CONN_PORT)
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}
	defer conn.Close()

	msg, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println("Read message failed:", err.Error())
	}

	conn.Write([]byte(msg))
}
