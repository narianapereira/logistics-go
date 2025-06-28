package ports

type ParserInterface interface {
	Parse(content []byte) ([]byte, error)
}
