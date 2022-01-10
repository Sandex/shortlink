package storage

// Интерфейс для хранения ссылок
type UrlStorage interface {
	// Привязать ссылку к хэшу
	Bind(url string, hash string)

	// Получить ссылку по хэшу
	Fetch(hash string) string
}
