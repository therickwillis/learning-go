package server

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var (
	eventStreams = map[chan string]bool{}
	streamsLock  sync.Mutex
)

func BroadcastUpdate() {
	streamsLock.Lock()
	defer streamsLock.Unlock()

	for ch := range eventStreams {
		select {
		case ch <- "ping":
		default:
		}
	}
}

func handleEvents(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[GET /events] Received SSE Request")
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	eventChan := make(chan string)
	streamsLock.Lock()
	eventStreams[eventChan] = true
	streamsLock.Unlock()

	defer func() {
		streamsLock.Lock()
		delete(eventStreams, eventChan)
		streamsLock.Unlock()
		close(eventChan)
	}()

	go func() {
		ticker := time.NewTicker(30 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-r.Context().Done():
				return
			case <-ticker.C:
				fmt.Fprintf(w, ": keep-alive\n\n")
				flusher.Flush()
			}
		}
	}()

	for msg := range eventChan {
		fmt.Println("[SSE] sending event:", msg)
		fmt.Fprintf(w, "event: update\n")
		fmt.Fprintf(w, "data: ping\n\n")
		flusher.Flush()
	}
}
