package storage

// URLStorage Интерфейс для хранения ссылок
type URLStorage interface {
	// Bind Привязать ссылку к хэшу
	Bind(url string, hash string)

	// Fetch Получить ссылку по хэшу
	Fetch(hash string) string
}
