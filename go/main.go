package main

import (
	"fmt"
	"os"
)

// Get command line arguments
func printArgs() {
	args := os.Args
	for i, v := range args {
		fmt.Println(i, "->", v)
	}
}

func main() {
}
