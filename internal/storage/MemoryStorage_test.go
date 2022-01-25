package storage

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFetch(t *testing.T) {
	s := new(MemoryStorage)
	s.Init()

	url := "test_url_1"
	hash := "test_hash_1"
	s.Bind(url, hash)

	//if tHash := s.fetch(hash); tHash != url {
	//	t.Errorf("For hash %s, expected ULR %s, got %s", hash, url, tHash)
	//}

	tHash := s.Fetch(hash)
	assert.Equal(t, tHash, url, "значения должны быть одинаковыми")
}
