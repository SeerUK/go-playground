package main

import (
	"fmt"
	"log"
)

func main() {
	var name string

	fmt.Print("Enter your name: ")
	if _, err := fmt.Scan(&name); err != nil {
		log.Fatalf("Unable to read name: %v", err)
	}

	fmt.Printf("Hello, %s!\n", name)
}
