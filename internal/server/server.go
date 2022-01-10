package server

import (
	"fmt"
	"github.com/Sandex/shortlink/internal/generator"
	"github.com/Sandex/shortlink/internal/handlers"
	"github.com/Sandex/shortlink/internal/storage"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

type ShortenerServer struct {
	storage   storage.UrlStorage
	generator generator.HasGenrator
}

// Start Запустить сервер
func (s *ShortenerServer) Start(addr string, storage storage.UrlStorage, generator generator.HasGenrator) {
	s.storage = storage
	s.generator = generator

	r := s.NewRouter()

	err := http.ListenAndServe(addr, r)
	if err != nil {
		fmt.Println("Can not start server ")
		fmt.Println(err)
	}
}

func (s *ShortenerServer) NewRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Post("/", func(res http.ResponseWriter, req *http.Request) {
		handlers.MakeShortHandler(res, req, s.generator, s.storage)
	})

	r.Get("/{hash}", func(res http.ResponseWriter, req *http.Request) {
		handlers.FetchUrlHandler(res, req, s.storage)
	})

	return r
}
