# syntax=docker/dockerfile:1
FROM golang:1.21.5 AS build-stage

WORKDIR /app

COPY go.mod ./
RUN go mod download


COPY templates templates/

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /api ./cmd/api

FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /app

COPY --from=build-stage /api /api
COPY --from=build-stage /app/templates templates/

EXPOSE 8080

USER nonroot:nonroot
CMD ["/api"]

