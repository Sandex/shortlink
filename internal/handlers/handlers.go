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

// FetchUrlHandler Метод сервера, возвращает Location в заголовке ответа для найденного хэша
func FetchUrlHandler(res http.ResponseWriter, req *http.Request, storage storage.UrlStorage) {
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
	} else {
		// send location
		log.Printf("URL: %s\n", urlOriginal)

		res.Header().Add("Location", urlOriginal)
		res.WriteHeader(http.StatusTemporaryRedirect)
	}
}

// MakeShortHandler Метод сервера, создает хэш для ссылки и сохраняет эту связку
func MakeShortHandler(res http.ResponseWriter, req *http.Request, generator generator.HasGenrator, storage storage.UrlStorage) {
	inputRawUrl, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	// validate URL
	urlStr := validateUrl(err, string(inputRawUrl))

	if urlStr == "" {
		log.Printf("Invalide URL\n")
		res.WriteHeader(http.StatusBadRequest)
	} else {
		log.Printf("Got URL: %s\n", urlStr)

		// do short and store
		hash := generator.MakeUrlId(urlStr)
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
		_, err = fmt.Fprintf(res, newLink.String())
		if err != nil {
			log.Printf("Error: %s\n", err)
		}
	}

}

func validateUrl(err error, inputRawUrl string) string {
	urlStr := ""
	inputUrl, err := url.Parse(inputRawUrl)
	if err == nil {
		urlStr = inputUrl.String()
	}

	return urlStr
}
