# syntax=docker/dockerfile:1

FROM golang:alpine as build
RUN apk add --no-cache build-base brotli-dev

FROM build

WORKDIR /app
COPY *.go ./
COPY go.mod ./

RUN CGO_LDFLAGS="-lbrotlienc -lbrotlidec -lbrotlicommon " CGO_ENABLED=1 GOOS=linux go test
