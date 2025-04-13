package server

import (
	"fmt"
	"net/http"
	"sync"
)

var (
	clients   = make(map[chan string]bool)
	clientsMu sync.Mutex
)

func handleEvents(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[Handle:Events] Request Received")

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	msgChan := make(chan string)
	clientsMu.Lock()
	clients[msgChan] = true
	clientsMu.Unlock()

	defer func() {
		clientsMu.Lock()
		delete(clients, msgChan)
		clientsMu.Unlock()
		close(msgChan)
	}()

	for {
		select {
		case <-r.Context().Done():
			clientsMu.Lock()
			defer clientsMu.Unlock()
			delete(clients, msgChan)
			return
		case msg := <-msgChan:
			fmt.Println("[SSE] sending event", msg)
			fmt.Fprintf(w, "event: update\n")
			fmt.Fprintf(w, "data: %v\n\n", msg)
			flusher.Flush()
		}
	}

}

func BroadcastGame() {
	message := "ping"
	clientsMu.Lock()
	defer clientsMu.Unlock()
	for ch := range clients {
		select {
		case ch <- message:
		default:
			// skip clients with full data channels
		}
	}

}
