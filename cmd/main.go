package main

import (
	"io"
	"log"
	"log/slog"
	"os"

	v1 "github.com/narianapereira/logistics-go/internal/adapter/http/v1"
	"github.com/narianapereira/logistics-go/internal/adapter/parser"
	"github.com/narianapereira/logistics-go/internal/application/service"
)

func main() {

	logHandler := slog.NewJSONHandler(io.Writer(os.Stdout), &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})
	logger := slog.New(logHandler)
	parserAdapter := parser.NewTextParser(logger)
	parserService := service.NewParserService(parserAdapter)
	router := v1.NewRouter(parserService)

	logger.Info("Server up and running on port 8080")
	err := router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
