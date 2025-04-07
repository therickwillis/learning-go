package main

import (
	"flag"
	"fmt"
)

func main() {
	start := flag.Int("start", 1, "The number to start at")
	end := flag.Int("end", 100, "The number to end at (inclusive)")

	for i := *start; i < *end; i++ {
		if i%15 == 0 {
			fmt.Println("Nibb High Football Rules")
		} else if i%3 == 0 {
			fmt.Println("Foo")
		} else if i%5 == 0 {
			fmt.Println("Bar")
		} else {
			fmt.Println(i)
		}
	}
}
