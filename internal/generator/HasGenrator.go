package generator

type HasGenrator interface {
	MakeURLID(urlStr string) string
}
