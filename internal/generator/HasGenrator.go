package generator

type HasGenrator interface {
	MakeUrlId(urlStr string) string
}
