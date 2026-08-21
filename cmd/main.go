package main

import (
	"log"

	"github.com/sthiroshima/go-cache/internal"
)

func main() {
	server := internal.NewServer("tcp", 6379)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
