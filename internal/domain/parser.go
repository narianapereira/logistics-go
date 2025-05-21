package domain

type ParserInterface interface {
	Parse(content []byte) (interface{}, error)
}
