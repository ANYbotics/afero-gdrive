# Preparing the build environment
FROM golang:1.16-alpine AS builder
ENV GOFLAGS="-mod=readonly"
RUN apk add --update --no-cache bash ca-certificates curl git
RUN mkdir -p /workspace
WORKDIR /workspace

# Building
COPY . .
RUN go build -v -o afero-gdrive
