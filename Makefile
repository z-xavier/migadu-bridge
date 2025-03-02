.PHONY: build-go build-frontend build-all clean tool help

all: build

# build 包括前端
build: clean
	npm ci --prefix frontend && \
	npm run build --prefix frontend && \
	cp -r frontend/dist internal/migadu-bridge/static && \
	go mod tidy && \
	go vet ./... && \
	go build -v -ldflags '-s -w' ./

build-frontend: clean
	npm ci --prefix frontend && \
	npm run build --prefix frontend

# build 只包括后端
build-go: clean
	go mod tidy && \
	go vet ./... && \
	go build -v -tags dev ./

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
	rm -rf frontend/dist
	rm -rf migadu-bridge

help:
	@echo "make: compile packages and dependencies"
	@echo "make tool: run specified go tool"
	@echo "make lint: golint ./..."
	@echo "make clean: remove object files and cached files"