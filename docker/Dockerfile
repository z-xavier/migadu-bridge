# Build stage

# Javascript build stage

# Golang build stage
FROM docker.io/library/golang:1.24-alpine3.21 AS go-builder

ARG APK_REPO
RUN if [ -n "${APK_REPO}" ]; then sed -i "s/dl-cdn.alpinelinux.org/${APK_REPO}/g" /etc/apk/repositories; fi

ARG GOPROXY
ENV GOPROXY=${GOPROXY:-direct}

# Build deps
RUN apk --no-cache add \
        git \
        && rm -rf /var/cache/apk/*

# Setup repo
COPY . ${GOPATH}/src/git.zxavier.com/xavier/migadu-bridge
WORKDIR ${GOPATH}/src/git.zxavier.com/xavier/migadu-bridge

# Begin env-to-ini build
ENV CGO_ENABLED=0

#RUN go mod tidy && \
RUN go vet ./... && \
    go build -v -ldflags '-s -w' ./

FROM docker.io/library/alpine:3.21

ARG APK_REPO
RUN if [ -n "${APK_REPO}" ]; then sed -i "s/dl-cdn.alpinelinux.org/${APK_REPO}/g" /etc/apk/repositories; fi

LABEL maintainer="me@zxavier.com"

EXPOSE 8080 8081

# Build deps
RUN apk --no-cache add \
        ca-certificates \
        && rm -rf /var/cache/apk/*

VOLUME ["/data", "/config"]

COPY --from=go-builder /go/src/git.zxavier.com/xavier/migadu-bridge/conf /config
COPY --from=go-builder /go/src/git.zxavier.com/xavier/migadu-bridge/migadu-bridge ./migadu-bridge

ENTRYPOINT ["./migadu-bridge/migadu-bridge"]
CMD ["-c", "/config/conf.yaml"]