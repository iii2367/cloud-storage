FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o cloud-storage ./cmd/app

FROM alpine:3.20

WORKDIR /app

COPY --from=builder /app/cloud-storage .
COPY --from=builder /app/web ./web

CMD ["./cloud-storage"]
