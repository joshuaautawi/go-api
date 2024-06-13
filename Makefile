.PHONY: all build run clean

all: build

build:
	go build -o bin/main cmd/main.go

run:
	go run cmd/main.go

clean:
	rm -rf bin/main