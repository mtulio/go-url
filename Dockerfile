FROM golang:1.10
WORKDIR /build
ADD *.go ./
RUN mkdir -p bin/
RUN CGO_ENABLED=0 GOOS=linux \
    go build -a -installsuffix cgo \
    -o ./bin/go-url *.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /go
COPY --from=0 /build/bin/go-url /go/go-url
ADD hack/docker-entrypoint.sh /go/entrypoint.sh
RUN chmod +x /go/entrypoint.sh
ENTRYPOINT ["sh","/go/entrypoint.sh"]
CMD [ "-h" ]