FROM golang:1.20.6-alpine3.18 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN  go mod download

COPY cmd cmd/
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -f arundas285/api -v -o ./ecommerce cmd/api

# final stage
FROM gcr.io/distroless/base-debian11 AS build-release-stage


WORKDIR /app

