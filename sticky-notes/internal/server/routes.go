package server

import (
	"fmt"
	"net/http"
)

func Run() error {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/events", handleEvents)
	http.HandleFunc("/poll", handlePoll)
	http.HandleFunc("/add", handleAdd)
	http.HandleFunc("/delete/", handleDelete)

	fmt.Println("Server Listening on http://localhost:8080")
	return http.ListenAndServe(":8080", nil)
}
