FROM golang:v1.24

WORKDIR $GOPATH/src/git.zxavier.com/xavier/migadu-bridge
COPY . $GOPATH/src/git.zxavier.com/xavier/migadu-bridge
RUN go build .

EXPOSE 8000
ENTRYPOINT ["./migadu-bridge"]
CMD ["-c", "/etc/migadu-bridge/conf.yaml"]