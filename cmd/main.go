package main

import (
	"log"

	"github.com/sthiroshima/go-cache/internal"
)

func main() {
	cfg, err := internal.CfgLoad()
	if err != nil {
		log.Fatal("Config validation error: ", err)
	}

	server := internal.NewServer("tcp", cfg.Port)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
