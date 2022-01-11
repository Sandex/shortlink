package handlers

import (
	"bytes"
	"github.com/go-chi/chi/v5"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Создаем mock хранилища
type StorageMock struct{}

func (s *StorageMock) Bind(url string, hash string) {}
func (s *StorageMock) Fetch(hash string) string {
	switch hash {
	case "YA":
		return "https://ya.ru"
	case "BING":
		return "https://bing.com"
	default:
		return ""
	}
}

// Создаем mock
type GeneratorMock struct{}

func (h *GeneratorMock) MakeUrlId(url string) string {
	switch url {
	case "https://ya.ru":
		return "YA"
	case "https://bing.com":
		return "BING"
	default:
		return ""
	}
}

func TestFetchUrlHandler(t *testing.T) {
	storageMock := new(StorageMock)

	// определяем структуру теста
	type want struct {
		code     int
		location string
	}
	// создаём массив тестов: имя и желаемый результат
	tests := []struct {
		name string
		url  string
		want want
	}{
		// определяем все тесты
		{
			name: "positive test #1",
			url:  "http://localhost:8080/YA",
			want: want{
				code:     307,
				location: "https://ya.ru",
			},
		},
		{
			name: "positive test #2",
			url:  "http://localhost:8080/BING",
			want: want{
				code:     307,
				location: "https://bing.com",
			},
		},
		{
			name: "negative test #3",
			url:  "http://localhost:8080/NO_EXISTS_HASH",
			want: want{
				code:     400,
				location: "",
			},
		},
		{
			name: "negative test #4",
			url:  "http://localhost:8080/",
			want: want{
				code:     404,
				location: "",
			},
		},
	}
	for _, tt := range tests {
		// запускаем каждый тест
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, tt.url, nil)

			// создаём новый Recorder
			w := httptest.NewRecorder()

			r := chi.NewRouter()
			r.Get("/{hash:[a-zA-Z0-9-]+}", func(res http.ResponseWriter, req *http.Request) {
				FetchUrlHandler(res, req, storageMock)
			})

			// запускаем сервер
			r.ServeHTTP(w, request)
			resp := w.Result()

			// проверяем код ответа
			if resp.StatusCode != tt.want.code {
				t.Errorf("Expected status code %d, got %d", tt.want.code, resp.StatusCode)
			}

			// заголовок ответа
			if resp.Header.Get("Location") != tt.want.location {
				t.Errorf("Expected Location %s, got %s", tt.want.location, resp.Header.Get("Location"))
			}
		})
	}
}

func TestMakeShortHandler(t *testing.T) {
	storageMock := new(StorageMock)
	generatorMock := new(GeneratorMock)

	// определяем структуру теста
	type want struct {
		code     int
		response string
	}
	// создаём массив тестов: имя и желаемый результат
	tests := []struct {
		name string
		url  string
		want want
	}{
		// определяем все тесты
		{
			name: "positive test #11",
			url:  "https://ya.ru",
			want: want{
				code:     201,
				response: "http://localhost:8080/YA",
			},
		},
		{
			name: "positive test #12",
			url:  "https://bing.com",
			want: want{
				code:     201,
				response: "http://localhost:8080/BING",
			},
		},
		{
			name: "negative test #13",
			url:  "",
			want: want{
				code:     406,
				response: "",
			},
		},
	}
	for _, tt := range tests {
		// запускаем каждый тест
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/", bytes.NewBufferString(tt.url))

			// создаём новый Recorder
			w := httptest.NewRecorder()
			// определяем хендлер
			h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				MakeShortHandler(w, r, generatorMock, storageMock)
			})

			// запускаем сервер
			h.ServeHTTP(w, request)
			res := w.Result()

			// проверяем код ответа
			if res.StatusCode != tt.want.code {
				t.Errorf("Expected status code %d, got %d", tt.want.code, w.Code)
			}

			// получаем и проверяем тело запроса
			defer res.Body.Close()
			resBody, err := io.ReadAll(res.Body)
			if err != nil {
				t.Fatal(err)
			}
			if string(resBody) != tt.want.response {
				t.Errorf("Expected body %s, got %s", tt.want.response, w.Body.String())
			}
		})
	}
}
