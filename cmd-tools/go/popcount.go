package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func popCount(n uint64) int {
	return int(pc[byte(n>>(0*8))] +
		pc[byte(n>>(1*8))] +
		pc[byte(n>>(2*8))] +
		pc[byte(n>>(3*8))] +
		pc[byte(n>>(4*8))] +
		pc[byte(n>>(5*8))] +
		pc[byte(n>>(6*8))] +
		pc[byte(n>>(7*8))])
}

func readPopCount(args []string) {
	for _, l := range args {
		l = strings.TrimSpace(l)
		num, err := strconv.ParseUint(l, 10, 64)
		if err != nil {
			log.Fatalf("Invalid Input: %v", err)
		}
		fmt.Printf("%d - Pop Count: %d\n", num, popCount(num))
	}
}
