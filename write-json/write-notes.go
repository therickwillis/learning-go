package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Note struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func main() {
	notes := []Note{
		{Title: "Groceries", Body: "Buy milk and eggs"},
		{Title: "Work", Body: "Finish Go tuturial"},
		{Title: "Game Idea", Body: "Build a tiny CLI RPG"},
	}

	file, err := os.Create("notes.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(notes)
	if err != nil {
		panic(err)
	}

	fmt.Println("Notes saved to notes.json")
}
