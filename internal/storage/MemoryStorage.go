package storage

// Реализация хранения в памяти
type MemoryStorage struct {
	storage map[string]string
}

func (s *MemoryStorage) Init() {
	s.storage = make(map[string]string)
}

func (s *MemoryStorage) Bind(url string, hash string) {
	s.storage[hash] = url
}

func (s *MemoryStorage) Fetch(hash string) string {
	url := s.storage[hash]

	return url
}
