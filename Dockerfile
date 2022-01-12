FROM golang:1.17 as builder
WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./bin/go-url ./cmd/go-url/

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /go/src/app/bin/go-url /app/
ADD hack/docker-entrypoint.sh /app/entrypoint.sh
RUN chmod +x /app/entrypoint.sh
ENTRYPOINT ["sh","/app/entrypoint.sh"]
CMD [ "-h" ]
