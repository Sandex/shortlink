package generator

import (
	gonanoid "github.com/matoous/go-nanoid/v2"
)

type NanoIdHasGenerator struct{}

func (h *NanoIdHasGenerator) MakeUrlId(url string) string {
	hash, _ := gonanoid.New()

	return hash
}
