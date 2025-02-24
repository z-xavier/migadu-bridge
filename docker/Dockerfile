# Build stage
FROM docker.io/library/golang:1.24-alpine3.21 AS build-env
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

# Checkout version if set
#RUN if [ -n "${MB_VERSION}" ]; then git checkout "${MB_VERSION}"; fi \
# && make clean-all build

# Begin env-to-ini build
RUN go build .

FROM docker.io/library/alpine:3.21
ARG APK_REPO
RUN if [ -n "${APK_REPO}" ]; then sed -i "s/dl-cdn.alpinelinux.org/${APK_REPO}/g" /etc/apk/repositories; fi

LABEL maintainer="me@zxavier.com"

EXPOSE 8080
EXPOSE 8081

# Build deps
RUN apk --no-cache add \
        ca-certificates \
        && rm -rf /var/cache/apk/*

VOLUME ["/data"]

ENTRYPOINT ["./migadu-bridge/migadu-bridge"]
CMD ["-c", "/etc/migadu-bridge/conf.yaml"]

COPY --from=build-env /tmp/local /
COPY --from=build-env /go/src/git.zxavier.com/xavier/migadu-bridge ./migadu-bridge