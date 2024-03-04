build:
	go build -o bin/main cmd/main.go

run:
	go run cmd/main.go

test:
	go test -v github.com/buirec/Frameplay-ID-to-Word/server

all: build test run