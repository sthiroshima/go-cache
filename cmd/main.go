package main

import (
	"log"
	"os"
	"strconv"

	"github.com/sthiroshima/go-cache/internal"
)

func main() {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatal("Validation PORT error: ", err)
	}

	server := internal.NewServer("tcp", port)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
