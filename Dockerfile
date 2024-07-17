FROM golang:1.20-alpine as builder

WORKDIR /Rescounts-Task

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o server ./cmd/server

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/server .

CMD ["./server"]
