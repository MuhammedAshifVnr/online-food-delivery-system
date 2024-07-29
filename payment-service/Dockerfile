FROM golang:1.22.0-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o payment-service ./cmd/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=build /app/payment-service .
COPY --from=build /app/pkg/temp /root/pkg/temp

EXPOSE 8080

CMD ["./payment-service"]
