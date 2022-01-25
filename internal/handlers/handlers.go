package handlers

import (
	"fmt"
	"github.com/Sandex/shortlink/internal/generator"
	"github.com/Sandex/shortlink/internal/storage"
	"github.com/go-chi/chi/v5"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// FetchURLHandler Метод сервера, возвращает Location в заголовке ответа для найденного хэша
func FetchURLHandler(res http.ResponseWriter, req *http.Request, storage storage.URLStorage) {
	// get url hash
	urlHash := chi.URLParam(req, "hash")
	log.Printf("Got hash: %s\n", urlHash)

	// fetch url
	urlOriginal := storage.Fetch(urlHash)
	if urlOriginal == "" {
		// not found
		log.Printf("Can not find URL for hash: %s\n", urlHash)

		res.WriteHeader(http.StatusBadRequest)
		_, err := fmt.Fprintf(res, "Can not find URL for this hash!")
		if err != nil {
			log.Printf("Error: %s\n", err)
		}

		return
	}

	// send location
	log.Printf("URL: %s\n", urlOriginal)

	res.Header().Add("Location", urlOriginal)
	res.WriteHeader(http.StatusTemporaryRedirect)
}

// MakeShortHandler Метод сервера, создает хэш для ссылки и сохраняет эту связку
func MakeShortHandler(res http.ResponseWriter, req *http.Request, generator generator.HasGenrator, storage storage.URLStorage) {
	inputRawURL, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Printf("Can not read body data\n")
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	// validate URL
	inputURL, err := url.Parse(string(inputRawURL))
	if err != nil {
		log.Printf("Invalide URL\n")
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	urlStr := inputURL.String()
	log.Printf("Got URL: %s\n", urlStr)

	// do short and store
	hash := generator.MakeURLID(urlStr)
	log.Printf("Generate HASH %s for URL: %s\n", hash, urlStr)
	storage.Bind(urlStr, hash)

	// output
	res.WriteHeader(http.StatusCreated)

	// build new link
	newLink := url.URL{
		Scheme: "http",
		Host:   req.Host,
		Path:   hash,
	}

	// send to client
	_, err = res.Write([]byte(newLink.String()))
	if err != nil {
		log.Printf("Can not write http body\n")
	}
}
