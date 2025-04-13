package main

import (
	"clicktrainer/internal/server"
	"clicktrainer/internal/targets"
	"log"
)

func main() {
	targets.Add()

	if err := server.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
