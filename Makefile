.PHONY: build run clean tidy fmt

build:
	go build -o bin/server cmd/server/main.go
run:
	go run cmd/server/main.go
clean:
	rm -rf bin/*
tidy:
	go mod tidy
fmt:
	go fmt ./...
