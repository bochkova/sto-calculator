FROM golang:alpine AS builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/api/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /build/app .
EXPOSE 8080

CMD ["./app"]