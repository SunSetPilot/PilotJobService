FROM golang:1.20 as builder

ENV GO111MODULE=on

WORKDIR /app

COPY . .

RUN go build -o service.bin main.go

RUN mkdir output && cp service.bin output/ && cp -r etc output/

FROM ubuntu:latest

WORKDIR /app

COPY --from=builder /app/output .

RUN mkdir /app/logs

CMD ["/app/service.bin", "-f", "/app/etc/config_dev.yaml"]