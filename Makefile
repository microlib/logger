.PHONY: all test build clean

all: test build

build: 
	mkdir -p build
	go build -o build ./...

build-dev:
	mkdir -p build
	GOOS=linux go build -ldflags="-s -w" -o build ./...

test:
	go test -v -coverprofile=tests/results/cover.out ./...

cover:
	go tool cover -html=tests/results/cover.out -o tests/results/cover.html

clean:
	rm -rf build/*
	go clean ./...
