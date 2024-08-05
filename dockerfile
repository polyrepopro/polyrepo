FROM golang:1.22.0-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o /app/polyrepo

FROM alpine:latest

COPY --from=builder /app/polyrepo /usr/bin/polyrepo

ENTRYPOINT ["/usr/bin/polyrepo"]