.PHONY: build
build: cmd/main.go
	go build -o bin/main cmd/main.go