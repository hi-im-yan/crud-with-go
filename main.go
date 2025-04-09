package main

import (
	"fmt"
	"log"

	"github.com/hi-im-yan/crud-with-go/server"
)

func main() {
	server := server.NewServer("8080")

	fmt.Println("Starting server on port " + server.Port)

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
