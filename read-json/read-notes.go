package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Note struct {
	Title string `json:"text"`
	Body  string `json:"body"`
}

func main() {
	//load the file
	file, err := os.Open("notes.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	//decode
	var notes []Note
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&notes)
	if err != nil {
		panic(err)
	}

	//print
	for i, note := range notes {
		fmt.Printf("[%d] %s: %s\n", i+1, note.Title, note.Body)
	}
}
