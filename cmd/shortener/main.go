package main

import (
	"github.com/Sandex/shortlink/internal/generator"
	"github.com/Sandex/shortlink/internal/server"
	"github.com/Sandex/shortlink/internal/storage"
	"log"
)

func main() {
	log.Println("Start server")

	// Make storage
	URLStorage := new(storage.MemoryStorage)
	URLStorage.Init()

	// Make hash generator
	hashGenerator := new(generator.NanoIDHasGenerator)

	// Make server
	srv := new(server.ShortenerServer)
	srv.Start("127.0.0.1:8080", URLStorage, hashGenerator)
}
