package main

import (
	"fmt"
	"github.com/asirolly/grpctest/internal/src/handler"
	"github.com/asirolly/grpctest/internal/src/middleware"
	"log"
	"log/slog"
	"net/http"
	"os"
)

const port = "8080"

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	healthHandler := &handler.Health{
		Logger: logger,
	}
	helloHandler := &handler.Hello{
		Logger: logger,
	}
	headersHandler := &handler.Headers{
		Logger: logger,
	}

	mux := http.NewServeMux()

	mux.Handle("GET /health/{$}", healthHandler)
	mux.Handle("GET /hello/{$}", helloHandler)
	mux.Handle("GET /headers/{$}", headersHandler)

	logger.Info("Starting server...", slog.String("port", port))
	log.Fatal(
		http.ListenAndServe(
			fmt.Sprintf(":%s", port),
			middleware.Wrap(mux, logger),
		),
	)
}
