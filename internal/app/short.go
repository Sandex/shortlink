package app

import (
	"encoding/base32"
	"math/rand"
	"time"
)

func MakeUrlId() string {
	rand.Seed(time.Now().UnixNano())

	randomBytes := make([]byte, 8)
	rand.Read(randomBytes)
	hash := string(base32.StdEncoding.EncodeToString(randomBytes)[:8])

	return hash
}
