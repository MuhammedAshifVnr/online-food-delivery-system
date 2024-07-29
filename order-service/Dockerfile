FROM golang:1.22.0-alpine AS build

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o order-service ./cmd

FROM alpine:latest

WORKDIR /root/

COPY --from=build /app/order-service .

EXPOSE 8080

CMD ["./order-service"]
