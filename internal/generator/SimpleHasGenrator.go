package generator

import (
	"encoding/base32"
	"math/rand"
	"time"
)

type SimpleHasGenerator struct{}

func (h *SimpleHasGenerator) MakeURLID(url string) string {
	rand.Seed(time.Now().UnixNano())

	randomBytes := make([]byte, 8)
	rand.Read(randomBytes)
	hash := base32.StdEncoding.EncodeToString(randomBytes)[:8]

	return hash
}
