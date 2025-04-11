package server

import (
	"fmt"
	"html/template"
	"net/http"
	"stickynotes/internal/notes"
	"strconv"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[GET /] Rendering Index")
	tmpl := template.Must(template.ParseFiles("templates/index.html", "templates/note.html"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		fmt.Println("index template execution error:", err)
	}
}

func handlePoll(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[GET /poll] Poll Request")
	tmpl := template.Must(template.ParseFiles("templates/note.html"))
	err := tmpl.ExecuteTemplate(w, "noteList", notes.Notes)
	if err != nil {
		fmt.Println("Poll template error:", err)
	}
}

func handleAdd(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Println("Error parsing form", r.URL.Path)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	content := r.FormValue("content")
	color := r.FormValue("color")

	fmt.Println("Adding Note:", content, color)
	notes.Add(content, color)
	BroadcastUpdate()
	w.WriteHeader(http.StatusOK)
}

func handleDelete(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[DELETE /delete/] Delete Request", r.URL.Path)

	idStr := r.URL.Path[len("/delete/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid ID")
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	notes.Delete(id)
	BroadcastUpdate()
	w.WriteHeader(http.StatusOK)
}
