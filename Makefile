.PHONY: build clean tool lint help

all: build

build:
	go build -v .

tool:
	@echo "=== Running vet ==="
	go vet ./... 2>&1 || true
	@echo "=== Formatting code ==="
	gofmt -l .

lint:
	golint ./...

clean:
	rm -rf go-gin-example
	go clean -i .

help:
	@echo "make: compile packages and dependencies"
	@echo "make tool: run specified go tool"
	@echo "make lint: golint ./..."
	@echo "make clean: remove object files and cached files"