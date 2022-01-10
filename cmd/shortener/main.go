package main

import (
	"fmt"
	"github.com/Sandex/shortlink/internal/generator"
	"github.com/Sandex/shortlink/internal/server"
	"github.com/Sandex/shortlink/internal/storage"
)

func main() {
	fmt.Println("Start server")

	// Make storage
	urlStorage := new(storage.MemoryStorage)
	urlStorage.Init()

	// Make hash generator
	hashGenerator := new(generator.NanoIdHasGenerator)

	// Make server
	srv := new(server.ShortenerServer)
	srv.Start("127.0.0.1:8080", urlStorage, hashGenerator)
}
