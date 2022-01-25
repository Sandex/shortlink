package storage

import (
	"fmt"
	"log"
)

// Реализация хранения в файле
type FileStorage struct {
	producer *Producer
	storage  map[string]string
}

func (s *FileStorage) Init(filename string) {
	producer, err := NewProducer(filename)

	s.producer = producer
	if err != nil {
		log.Fatal(fmt.Sprint("Can not init filename %s", filename))
	}

	// init store
	s.storage = make(map[string]string)
	for {
		readedEvent, _ := s.producer.ReadEvent()
		log.Printf("Read from file: %s", readedEvent)

		if readedEvent == nil {
			break
		}

		s.storage[readedEvent.Hash] = readedEvent.URL
	}
}

func (s *FileStorage) Bind(url string, hash string) {
	s.storage[hash] = url

	event := Event{URL: url, Hash: hash}
	if err := s.producer.WriteEvent(&event); err != nil {
		log.Fatal(err)
	}
}

func (s *FileStorage) Fetch(hash string) string {
	url := s.storage[hash]

	return url
}
