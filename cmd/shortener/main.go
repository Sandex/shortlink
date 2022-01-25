package main

import (
	"flag"
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

	serverAddress := flag.String("a", "", "Server address and port, ex.: localhost:8080")
	baseURL := flag.String("b", "", "Base URL")
	fileStoragePath := flag.String("f", "", "File storage path")
	flag.Parse()

	if *serverAddress != "" {
		cfg.ServerAddress = *serverAddress
	}

	if *baseURL != "" {
		cfg.ServerAddress = *baseURL
	}
	if *fileStoragePath != "" {
		cfg.ServerAddress = *fileStoragePath
	}

	log.Printf("Use SERVER_ADDRESS %s", cfg.ServerAddress)
	log.Printf("Use BASE_URL %s", cfg.BaseURL)
	log.Printf("Use FILE_STORAGE_PATH %s", cfg.FileStoragePath)

	// Make storage
	URLStorage := new(storage.FileStorage)
	URLStorage.Init(cfg.FileStoragePath)

	// Make hash generator
	hashGenerator := new(generator.NanoIDHasGenerator)

	// Make server
	srv := new(server.ShortenerServer)
	srv.Start(cfg, URLStorage, hashGenerator)
}
