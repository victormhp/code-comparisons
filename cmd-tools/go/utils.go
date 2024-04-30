package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

// Removes any prefix of s that looks like a file system path with components separated by slashes,
// and it removes any suffix that looks like a file type
func basename(s string) {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	fmt.Println(s)
}

// Basic curl command implementation
func fetch(urls []string) {
	for _, url := range urls {
		res, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		b, err := io.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}

		fmt.Println(b)
		fmt.Println("Status: ", res.Status)
	}
}

// Compares execution time of several fetch calls
func fetchAll(urls []string) {
	start := time.Now()
	ch := make(chan string)
	for _, url := range urls {
		go func() {
			start := time.Now()
			res, err := http.Get(url)
			if err != nil {
				ch <- fmt.Sprint(err)
				return
			}

			nbytes, err := io.Copy(io.Discard, res.Body)
			res.Body.Close()
			if err != nil {
				ch <- fmt.Sprintf("while reading %s: %v\n", url, err)
			}
			secs := time.Since(start).Seconds()
			ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
		}()
	}

	for range urls {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

// Popcount utility (counts the number of 1 in a binary)
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

// Unix uniq command, looks for duplicate lines in the standard input or a given file.
func uniq(files []string) {
	counts := make(map[string]int)
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "uniq-go %v\n", err)
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
