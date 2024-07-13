package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"regexp"
)

const port = `:8090`

type basicHandler struct {
	logger *slog.Logger
}

type badHandler struct {
	logger *slog.Logger
}

var (
	helloRe  = regexp.MustCompile(`^/hello/?$`)
	headerRe = regexp.MustCompile(`^/headers/?$`)
)

func (h *badHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.logger.Info(fmt.Sprintf("Endpoint does not exist. %s", r.URL.Path))
	http.Error(
		w,
		fmt.Sprintf(
			"%d %s",
			http.StatusNotFound,
			http.StatusText(http.StatusNotFound),
		), http.StatusNotFound,
	)
}

func (h *basicHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodGet && helloRe.MatchString(r.URL.Path):
		h.getHost(w)
	case r.Method == http.MethodGet && headerRe.MatchString(r.URL.Path):
		h.getHeaders(w, r)
	default:
		h.logger.Info("User supplied bad endpoint.")
		http.Error(w, fmt.Sprintf("Endpoint path not handled. %s", r.URL.Path), http.StatusBadRequest)
	}
}

func (h *basicHandler) getHost(w http.ResponseWriter) {
	hn := getHostName()
	h.logger.Info("getHost() called.")
	_, err := fmt.Fprintf(w, "Hello World!\nFrom %s", hn)
	if err != nil {
		h.logger.Error("Error writing response.", err)
	}
}

func (h *basicHandler) getHeaders(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("getHeaders() called.")
	for name, headers := range r.Header {
		for _, hdr := range headers {
			_, err := fmt.Fprintf(w, "%s: %s\n", name, hdr)
			if err != nil {
				h.logger.Error("Error writing response.", err)
			}
		}
	}
}

func getHostName() string {
	hn, _ := os.Hostname()
	return hn
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	loggedHandler := &basicHandler{
		logger,
	}

	loggedBadHandler := &badHandler{
		logger,
	}

	mux := http.NewServeMux()

	mux.Handle("/hello", loggedHandler)
	mux.Handle("/hello/", loggedHandler)
	mux.Handle("/headers", loggedHandler)
	mux.Handle("/headers/", loggedHandler)
	mux.Handle("/*", loggedBadHandler)

	logger.Info(fmt.Sprintf("Listening on port %s...", port))
	err := http.ListenAndServe(port, mux)
	if err != nil {
		logger.Error("Error starting webserver.", err)
	}
}
