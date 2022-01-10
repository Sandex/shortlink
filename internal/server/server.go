package server

import (
	"fmt"
	"github.com/Sandex/shortlink/internal/handlers"
	"github.com/Sandex/shortlink/internal/storage"
	"net/http"
)

type ShortenerServer struct {
	storage storage.UrlStorage
}

// Запустить сервер
func (s *ShortenerServer) Start(addr string, storage storage.UrlStorage) {
	s.storage = storage

	mux := http.NewServeMux()
	mux.HandleFunc("/", s.handle)

	server := http.Server{
		Addr:    addr,
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Can not start server ")
		fmt.Println(err)
	}
}

// Обработчик запросов
func (s *ShortenerServer) handle(res http.ResponseWriter, req *http.Request) {
	switch req.Method {

	case http.MethodPost:
		handlers.MakeShortHandler(res, req, s.storage)

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
