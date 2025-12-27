FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o authApp

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/authApp .

CMD ["/app/authApp"]