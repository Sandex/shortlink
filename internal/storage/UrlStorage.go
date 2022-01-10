package storage

// UrlStorage Интерфейс для хранения ссылок
type UrlStorage interface {
	// Bind Привязать ссылку к хэшу
	Bind(url string, hash string)

	// Fetch Получить ссылку по хэшу
	Fetch(hash string) string
}
