package server

import (
	"fmt"
	"net/http"
)

func handleEvents(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[Handle:Events] Request Received")

}
