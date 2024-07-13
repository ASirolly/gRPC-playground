package main

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"time"
)

const host = "localhost"
const port = 8090
const path = "hello/"

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	client := http.Client{
		Timeout: 5 * time.Second,
	}

	reqURL := fmt.Sprintf("http://%s:%d/%s", host, port, path)
	req, err := http.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		logger.Error("Could not construct request.", "dump", err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		logger.Error("Did not get a valid response from server.", "dump", err)
		return
	}

	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Error("Could not read response body.", "dump", err)
		return
	}

	logger.Info(string(bodyBytes), "status", resp.StatusCode)
}
