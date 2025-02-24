.PHONY: build build-all clean tool help

all: build

# build 只包括后端
build: clean
	go mod tidy && \
	go vet ./... && \
	go build -v ./

# build 包括前端
build-all: clean
	go mod tidy && \
	go vet ./... && \
	go build -v -ldflags '-s -w' ./

build-docker:
	cp docker/Dockerfile ./Dockerfile
	docker build -t migadu-bridge .
	rm ./Dockerfile

tool:
	@echo "=== Running vet ==="
	go vet ./... 2>&1 || true
	@echo "=== Formatting code ==="
	gofmt -l .

clean:
	rm -rf migadu-bridge

help:
	@echo "make: compile packages and dependencies"
	@echo "make tool: run specified go tool"
	@echo "make lint: golint ./..."
	@echo "make clean: remove object files and cached files"