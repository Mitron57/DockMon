FROM golang:1.23 AS builder

WORKDIR backend

COPY . .

RUN go mod download && go mod verify
RUN go build -v -o /usr/bin/backend ./cmd/monitor/main.go

FROM ubuntu:25.04

WORKDIR app

COPY config/config.yaml config.yaml
COPY .env .env
COPY --from=builder /usr/bin/backend ./backend

CMD ["./backend", "-c", "./config.yaml"]