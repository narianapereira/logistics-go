package main

import (
	v1 "github.com/narianapereira/logistics-go/internal/adapter/http/v1"
	"github.com/narianapereira/logistics-go/internal/adapter/parser"
	"github.com/narianapereira/logistics-go/internal/application/service"
	"log"
)

func main() {

	parserAdapter := parser.NewTextParser()
	parserService := service.NewParserService(parserAdapter)
	router := v1.NewRouter(parserService)

	log.Println("Server up and running")
	err := router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
