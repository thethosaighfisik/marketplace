FROM golang:1.24-alpine AS builder

RUN mkdir -p /home/thethosaighfisic/marketplace/cmd/app/config

WORKDIR /home/thethosaighfisic/marketplace

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .


RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main ./cmd/app 

FROM alpine:latest

RUN mkdir -p /home/thethosaighfisic/marketplace/cmd/app/config

COPY --from=builder /app/main /home/thethosaighfisic/marketplace/main
COPY --from=builder /home/thethosaighfisic/marketplace/cmd/app/config/config.yaml /home/thethosaighfisic/marketplace/cmd/app/config/

EXPOSE 8090
WORKDIR /home/thethosaighfisic/marketplace

CMD ["./main"]
