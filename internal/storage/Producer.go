package storage

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
)

type Event struct {
	URL  string `json:"url"`
	Hash string `json:"hash"`
}

type Producer struct {
	file    *os.File
	writer  *bufio.Writer
	scanner *bufio.Scanner
}

func NewProducer(filename string) (*Producer, error) {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		return nil, err
	}

	return &Producer{
		file:    file,
		writer:  bufio.NewWriter(file),
		scanner: bufio.NewScanner(file),
	}, nil
}

func (p *Producer) WriteEvent(event *Event) error {
	data, err := json.Marshal(&event)
	if err != nil {
		return err
	}

	if _, err := p.writer.Write(data); err != nil {
		return err
	}

	if err := p.writer.WriteByte('\n'); err != nil {
		return err
	}

	return p.writer.Flush()
}

func (p *Producer) ReadEvent() (*Event, error) {
	if !p.scanner.Scan() {
		return nil, p.scanner.Err()
	}
	data := p.scanner.Bytes()

	log.Println(data)

	event := Event{}
	err := json.Unmarshal(data, &event)
	if err != nil {
		return nil, err
	}

	return &event, nil
}
