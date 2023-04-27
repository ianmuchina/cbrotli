# syntax=docker/dockerfile:1

FROM golang:alpine as build
RUN apk add --no-cache build-base git cmake make bash
RUN cd /usr/local && git clone --depth=1 https://github.com/google/brotli && cd brotli && mkdir out && cd out && ../configure-cmake --disable-debug && make && make install

FROM build

WORKDIR /app
COPY . ./

RUN CGO_ENABLED=1 GOOS=linux go test
