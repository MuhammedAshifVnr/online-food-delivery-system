FROM golang:1.22.0-alpine AS build

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o user-service ./cmd

FROM alpine:latest

WORKDIR /root/

COPY --from=build /app/user-service .

EXPOSE 8080

CMD ["./user-service"]
