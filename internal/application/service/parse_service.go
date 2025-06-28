package service

import ports "github.com/narianapereira/logistics-go/internal/application/ports"

type ParserService struct {
	parser ports.ParserInterface
}

func NewParserService(parser ports.ParserInterface) *ParserService {
	return &ParserService{
		parser: parser,
	}
}

func (s *ParserService) Parse(content []byte) ([]byte, error) {
	return s.parser.Parse(content)
}
