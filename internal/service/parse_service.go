package service

import "github.com/narianapereira/logistics-go/internal/domain"

type ParserService struct {
	parser domain.ParserInterface
}

func NewParserService(parser domain.ParserInterface) *ParserService {
	return &ParserService{
		parser: parser,
	}
}

func (s *ParserService) Parse(content []byte) (interface{}, error) {
	return s.parser.Parse(content)
}
