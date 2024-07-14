package handler

import (
	"fmt"
	"log/slog"
	"net/http"
)

type Headers struct {
	Logger *slog.Logger
}

func (h *Headers) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	hdrmap := getHeaders(r)
	for name, hdr := range hdrmap {
		fmt.Fprintf(w, "%s: %s\n", name, hdr)
	}
}

func getHeaders(r *http.Request) map[string]string {
	hdrMap := map[string]string{}
	for name, headers := range r.Header {
		for _, hdr := range headers {
			hdrMap[name] = hdr
		}
	}
	return hdrMap
}
