FROM golang:1.10 AS build-env
WORKDIR /go
ADD *.go /go/
RUN mkdir -p /go/bin/
RUN go build -o ./bin/go-url *.go

#TODO fix multi-staging builder
# FROM alpine:latest
# RUN apk --no-cache add ca-certificates
# WORKDIR /go
# COPY --from=0 /go/bin/go-url /go/go-url
ADD scripts/docker-entrypoint.sh /go/entrypoint.sh
ENTRYPOINT ["/go/entrypoint.sh"]
CMD [ "-h" ]