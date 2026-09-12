FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -v -a -installsuffix cgo -o finance ./cmd/server

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/finance .

EXPOSE 8080

CMD ["./finance"]