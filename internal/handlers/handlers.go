package handlers

import (
	"fmt"
	"github.com/Sandex/shortlink/internal/app"
	"github.com/Sandex/shortlink/internal/storage"
	"io/ioutil"
	"net/http"
	"strings"
)

// Метод сервера, возвращает Location в заголовке ответа для найденного хэша
func FetchUrlHandler(res http.ResponseWriter, req *http.Request, storage storage.UrlStorage) {
	// get url hash
	urlHash := strings.TrimPrefix(req.URL.Path, "/")
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

// Метод сервера, создает хэш для ссылки и сохраняет эту связку
func MakeShortHandler(res http.ResponseWriter, req *http.Request, storage storage.UrlStorage) {
	url, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	// sanitize URL
	urlStr := strings.Trim(string(url), "\n\r\t ")
	fmt.Printf("Got url: %s\n", urlStr)

	// do short and store
	shortUrl := app.MakeUrlId()
	fmt.Printf("Generate HASH %s for URL: %s\n", shortUrl, urlStr)
	storage.Bind(urlStr, shortUrl)

	// output
	res.WriteHeader(http.StatusCreated)
	_, err = fmt.Fprintf(res, "http://"+req.Host+"/"+shortUrl)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}
