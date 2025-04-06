package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	message := flag.String("message", "hi", "Message to print")
	shout := flag.Bool("shout", false, "Whether to shout the message")
	repeat := flag.Int("repeat", 1, "How many times to repeat the message")

	flag.Parse()

	output := *message
	if *shout {
		output = strings.ToUpper(output)
	}

	for range *repeat {
		fmt.Println(output)
	}
}
