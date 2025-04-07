package main

import (
	"fmt"
	"net/http"
	"time"
)

func fetchStatus(url string, ch chan string) {
	start := time.Now()
	resp, err := http.Get(url)
	duration := time.Since(start)

	if err != nil {
		ch <- fmt.Sprintf("%s -> ERROR: %v", url, err)
		return
	}
	defer resp.Body.Close()

	ch <- fmt.Sprintf("%s -> %d (%v)", url, resp.StatusCode, duration)
}

func main() {
	urls := []string{
		"https://google.com",
		"https://github.com",
		"https://golang.org",
		"https://thisurldoesnotexist",
	}

	ch := make(chan string)

	for _, url := range urls {
		go fetchStatus(url, ch)
	}

	for range urls {
		fmt.Println(<-ch)
	}
}
