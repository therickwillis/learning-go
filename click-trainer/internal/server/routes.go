package server

import (
	"fmt"
	"net/http"
	"text/template"
)

var (
	tmpl *template.Template
)

func Run() error {
	tmpl = template.Must(template.ParseFiles("templates/game.html", "templates/join.html", "templates/target.html"))

	// new player flow (no cookie):
	//  "/" -  Set Name, Click Start -> Server generates a new game session
	//  "/g/?s=:sessionID"

	// new player flow (with cookie):
	// "/" - Wecome Back, name, Click Start -> server generates new game session
	// "/g/?s=:sessionID"

	// new player flow (no cookie, with session invite)
	// "/" - Set Name, Join
	// "/g/?s=:sessionID"

	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/g/", handleGame)
	http.HandleFunc("/register", handleRegister)
	http.HandleFunc("/target/", handleTarget)
	http.HandleFunc("/events", handleEvents)
	// http.HandleFunc("/poll", handlePoll)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Println("Server listening on http://localhost:8080")
	return http.ListenAndServe("0.0.0.0:8080", nil)
}
