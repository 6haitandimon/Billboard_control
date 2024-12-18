FROM golang:1.23.3 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o billboard main.go

FROM ubuntu:22.04

RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY --from=builder /app/billboard .

COPY ./ADS /app/ADS

COPY wait-for-it.sh /app/
RUN chmod +x /app/wait-for-it.sh

EXPOSE 8080

CMD ["/app/billboard"]
