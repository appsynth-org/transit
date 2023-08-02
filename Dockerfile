FROM golang:1.20 as base
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download

FROM base as builder
WORKDIR /app
COPY . ./
RUN go build -o app ./cmd/http/main.go

FROM debian:stable-slim as release
WORKDIR /app
COPY --from=builder /app/app ./
RUN apt update && apt install -y ca-certificates
EXPOSE 80
CMD ["./app"]