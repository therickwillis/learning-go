package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Note struct {
	ID      int
	Content string
	Color   string
}

var notes = []Note{}

var nextID = 1

func main() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/add", handleAdd)
	fmt.Println("Started serving on port 8080")
	http.ListenAndServe(":8080", nil)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html", "templates/note.html"))
	tmpl.Execute(w, notes)
}

func handleAdd(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	content := r.FormValue("content")
	color := r.FormValue("color")

	note := Note{
		ID:      nextID,
		Content: content,
		Color:   color,
	}
	nextID++

	notes = append(notes, note)

	if r.Header.Get("HX-Request") == "true" {
		tmpl := template.Must(template.ParseFiles("templates/note.html"))
		tmpl.Execute(w, note)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
