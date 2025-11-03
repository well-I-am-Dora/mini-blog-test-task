FROM golang:alpine AS builder

LABEL authors="darya"

WORKDIR /build

ADD go.mod .

COPY . .

RUN go build -o main cmd/main.go

CMD ["./main"]