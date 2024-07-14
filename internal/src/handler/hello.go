package handler

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

type Hello struct {
	Logger *slog.Logger
}

func (h *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from %s!", getHostName())
}

func getHostName() string {
	hn, _ := os.Hostname()
	return hn
}
