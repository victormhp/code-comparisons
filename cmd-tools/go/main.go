package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	CMD_UNIQ     = "uniq"
	CMD_FETCH    = "fetch"
	CMD_FETCHALL = "fetchall"
	CMD_POPCOUNT = "popcount"
)

func main() {
	cmd := flag.String("command", "", "Command Line Tool")
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Provide a command")
		os.Exit(1)
	}

	switch *cmd {
	case CMD_UNIQ:
		uniq(args)
	case CMD_FETCH:
		fetch(args)
	case CMD_FETCHALL:
		fetchAll(args)
	case CMD_POPCOUNT:
		readPopCount(args)
	default:
		fmt.Printf("Unknown tool: %s\n", *cmd)
		os.Exit(1)
	}
}
