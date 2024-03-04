build:
	go build -o bin/main cmd/main.go

run:
	go run cmd/main.go

lint:
	golangci-lint run

test:
	go test -v github.com/buirec/Frameplay-ID-to-Word/server

all: build lint test run