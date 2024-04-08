package main

import (
	"fmt"
	"os"
)

const (
	CMD_UNIQ = "-uniq"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Provide a unix command")
		os.Exit(1)
	}

	tool := args[0]
	flags := args[1:]
	switch tool {
	case CMD_UNIQ:
		uniq(flags)
	default:
		fmt.Printf("Unknown tool: %s\n", tool)
		os.Exit(1)
	}
}
