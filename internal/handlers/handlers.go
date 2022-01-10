package handlers

import (
	"fmt"
	"github.com/Sandex/shortlink/internal/generator"
	"github.com/Sandex/shortlink/internal/storage"
	"github.com/go-chi/chi/v5"
	"io/ioutil"
	"net/http"
	"strings"
)

// FetchUrlHandler Метод сервера, возвращает Location в заголовке ответа для найденного хэша
func FetchUrlHandler(res http.ResponseWriter, req *http.Request, storage storage.UrlStorage) {
	// get url hash
	urlHash := strings.TrimPrefix(chi.URLParam(req, "hash"), "/")
	fmt.Printf("Got hash: %s\n", urlHash)

	// fetch url
	urlOriginal := storage.Fetch(urlHash)
	if urlOriginal != "" {
		// send location
		fmt.Printf("URL: %s\n", urlOriginal)

		res.Header().Add("Location", urlOriginal)
		res.WriteHeader(http.StatusTemporaryRedirect)
	} else {
		// not found
		fmt.Printf("Can not find URL for hash: %s\n", urlHash)

		res.WriteHeader(http.StatusBadRequest)
		_, err := fmt.Fprintf(res, "Can not find URL for this hash!")
		if err != nil {
			fmt.Printf("Error: %s\n", err)
		}
	}
}

// MakeShortHandler Метод сервера, создает хэш для ссылки и сохраняет эту связку
func MakeShortHandler(res http.ResponseWriter, req *http.Request, generator generator.HasGenrator, storage storage.UrlStorage) {
	url, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	// sanitize URL
	urlStr := strings.Trim(string(url), "\n\r\t ")

	if urlStr != "" {
		fmt.Printf("Got url: %s\n", urlStr)

		// do short and store
		hash := generator.MakeUrlId(urlStr)
		fmt.Printf("Generate HASH %s for URL: %s\n", hash, urlStr)
		storage.Bind(urlStr, hash)

		// output
		res.WriteHeader(http.StatusCreated)
		_, err = fmt.Fprintf(res, "http://"+req.Host+"/"+hash)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
		}
	} else {
		fmt.Printf("Got empty url\n")
		res.WriteHeader(http.StatusNotAcceptable)
	}
}
