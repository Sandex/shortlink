package main

import (
	"fmt"
	"github.com/Sandex/shortlink/internal/server"
	"github.com/Sandex/shortlink/internal/storage"
)

func main() {
	fmt.Println("Start server")

	// Make storage
	urlStorage := new(storage.MemoryStorage)
	urlStorage.Init()

	// Make server
	srv := new(server.ShortenerServer)
	srv.Start("127.0.0.1:8080", urlStorage)
}
