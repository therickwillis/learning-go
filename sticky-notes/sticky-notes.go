package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
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
	http.HandleFunc("/delete/", handleDelete)
	fmt.Println("Started serving on port 8080")
	http.ListenAndServe(":8080", nil)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[GET /] Rendering Index")
	tmpl := template.Must(template.ParseFiles("templates/index.html", "templates/note.html"))
	err := tmpl.Execute(w, notes)
	if err != nil {
		fmt.Println("index template execution error:", err)
	}
}

func handleAdd(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[POST /add] Received form submission")

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	content := r.FormValue("content")
	color := r.FormValue("color")

	fmt.Println("Form Values", content, color)

	note := Note{
		ID:      nextID,
		Content: content,
		Color:   color,
	}
	nextID++

	notes = append(notes, note)

	if r.Header.Get("HX-Request") == "true" {
		fmt.Print("HX-Request Add")
		tmpl := template.Must(template.ParseFiles("templates/note.html"))
		err := tmpl.ExecuteTemplate(w, "note", note)
		if err != nil {
			fmt.Println("note template execution error:", err)
		}
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func handleDelete(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[DELETE /delete] Received delete request")

	idStr := r.URL.Path[len("/delete/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid ID")
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	fmt.Println("[DELETE /delete]", id)

	var updated []Note
	for _, n := range notes {
		if n.ID != id {
			updated = append(updated, n)
		}
	}
	notes = updated

	w.WriteHeader(http.StatusOK)
}
