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

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/", s.handle)
	r.Get("/", s.handle)
	err := http.ListenAndServe(addr, r)
	if err != nil {
		fmt.Println("Can not start server ")
		fmt.Println(err)
	}
}

// Обработчик запросов
func (s *ShortenerServer) handle(res http.ResponseWriter, req *http.Request) {
	switch req.Method {

	case http.MethodPost:
		handlers.MakeShortHandler(res, req, s.generator, s.storage)

	case http.MethodGet:
		handlers.FetchUrlHandler(res, req, s.storage)

	default:
		res.WriteHeader(http.StatusBadRequest)
		_, err := fmt.Fprintf(res, "Request must be POST or GET!")
		if err != nil {
			fmt.Printf("Error: %s\n", err)
		}
	}
}
