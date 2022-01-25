package generator

import (
	gonanoid "github.com/matoous/go-nanoid/v2"
)

type NanoIDHasGenerator struct{}

func (h *NanoIDHasGenerator) MakeURLID(url string) string {
	hash, _ := gonanoid.New()

	return hash
}
