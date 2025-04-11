package main

import (
	"log"
	"stickynotes/internal/server"
)

func main() {
	log.Default().Println("Starting Stickynotes Server")

	err := server.Run()
	if err != nil {
		log.Fatalln("Fatal error during startup", err)
	}
}
