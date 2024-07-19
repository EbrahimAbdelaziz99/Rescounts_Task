FROM golang:1.22.5-alpine AS builder

WORKDIR /Rescounts-Task

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o server ./cmd/server

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /Rescounts-Task/server .

COPY .env .env

CMD ["./server"]
