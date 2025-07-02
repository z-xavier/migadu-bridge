# Build stage
FROM docker.io/node:22-alpine AS node-builder

ARG NPM_REPO
RUN if [ -n "${NPM_REPO}" ]; then npm config set registry ${NPM_REPO}; fi

# Setup repo
COPY frontend /frontend

RUN cd /frontend && \
    npm ci && \
    npm run build

# Golang build stage
FROM docker.io/library/golang:1.24-alpine AS go-builder

ARG GOPROXY
ENV GOPROXY=${GOPROXY:-direct}

# Setup repo
COPY .. ${GOPATH}/src/git.zxavier.com/xavier/migadu-bridge
WORKDIR ${GOPATH}/src/git.zxavier.com/xavier/migadu-bridge

# Begin env-to-ini build
ENV CGO_ENABLED=0

COPY --from=node-builder /frontend/dist ${GOPATH}/src/git.zxavier.com/xavier/migadu-bridge/internal/migadubridge/static/dist

#RUN go mod tidy && \
RUN go vet ./... && \
    go build -v -ldflags '-s -w' ./

FROM docker.io/library/alpine

ARG APK_REPO
RUN if [ -n "${APK_REPO}" ]; then sed -i "s/dl-cdn.alpinelinux.org/${APK_REPO}/g" /etc/apk/repositories; fi

LABEL maintainer="me@zxavier.com"

EXPOSE 8080 8081

# Build deps
RUN apk --no-cache add \
        curl \
        ca-certificates \
        && rm -rf /var/cache/apk/*

VOLUME ["/data", "/config"]

COPY --from=go-builder /go/src/git.zxavier.com/xavier/migadu-bridge/conf /config
COPY --from=go-builder /go/src/git.zxavier.com/xavier/migadu-bridge/migadu-bridge ./migadu-bridge

ENTRYPOINT ["./migadu-bridge"]
CMD ["-c", "/config/conf.yaml"]
