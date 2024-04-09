package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

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

func fetchall(urls []string) {
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
