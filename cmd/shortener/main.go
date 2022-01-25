package main

import (
	"github.com/Sandex/shortlink/internal/config"
	"github.com/Sandex/shortlink/internal/generator"
	"github.com/Sandex/shortlink/internal/server"
	"github.com/Sandex/shortlink/internal/storage"
	"github.com/caarlos0/env/v6"
	"log"
)

func main() {
	log.Println("Start server")

	var cfg config.Config
	err := env.Parse(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Use SERVER_ADDRESS %s", cfg.ServerAddress)
	log.Printf("Use BASE_URL %s", cfg.BaseURL)

	// Make storage
	URLStorage := new(storage.MemoryStorage)
	URLStorage.Init()

	// Make hash generator
	hashGenerator := new(generator.NanoIDHasGenerator)

	// Make server
	srv := new(server.ShortenerServer)
	srv.Start(cfg, URLStorage, hashGenerator)
}
