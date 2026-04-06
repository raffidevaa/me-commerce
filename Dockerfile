FROM golang:1.24-alpine AS builder

WORKDIR /app  

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o me-commerce ./cmd/main.go

FROM alpine:3.21 AS final

WORKDIR /app  

COPY --from=builder /app/me-commerce .

EXPOSE 8080

CMD ["./me-commerce"]
