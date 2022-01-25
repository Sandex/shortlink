package server

import (
	"context"
	"github.com/Sandex/shortlink/internal/generator"
	"github.com/Sandex/shortlink/internal/handlers"
	"github.com/Sandex/shortlink/internal/storage"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type ShortenerServer struct {
	storage   storage.URLStorage
	generator generator.HasGenrator
}

// Start Запустить сервер
func (s *ShortenerServer) Start(addr string, storage storage.URLStorage, generator generator.HasGenrator) {
	s.storage = storage
	s.generator = generator

	r := s.NewRouter()

	srv := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Print("Server Started")

	<-done
	log.Print("Server Stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		// extra handling here
		cancel()
	}()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}
	log.Print("Server Exited Properly")
}

func (s *ShortenerServer) NewRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Post("/", func(res http.ResponseWriter, req *http.Request) {
		handlers.MakeShortHandler(res, req, s.generator, s.storage)
	})

	r.Get("/{hash:[A-Za-z0-9_-]+}", func(res http.ResponseWriter, req *http.Request) {
		handlers.FetchURLHandler(res, req, s.storage)
	})

	r.Post("/api/shorten", func(res http.ResponseWriter, req *http.Request) {
		handlers.APIShortenHandler(res, req, s.generator, s.storage)
	})

	return r
}
