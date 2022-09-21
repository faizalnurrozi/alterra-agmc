# step 1: build executable binary
FROM golang:1.18-alpine AS builder
LABEL maintainer="Faizal Nur Rozi<faizalnurrozi@gmail.com>"
RUN apk update && apk add --no-cache git && apk add --no-cach bash && apk add build-base
WORKDIR /app
COPY ../../ .
RUN go mod download
COPY docker/go .
RUN go build -o /agmc-service

# step 2: build a small image
FROM alpine:3.16.0
WORKDIR /app
COPY --from=builder agmc-service .
EXPOSE 8080
CMD ["./agmc-service", "-m=migrate"]