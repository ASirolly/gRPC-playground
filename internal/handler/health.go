package handler

import (
	"fmt"
	"log/slog"
	"net/http"
)

type Health struct {
	Logger *slog.Logger
}

func (h *Health) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Healthy.")
}
