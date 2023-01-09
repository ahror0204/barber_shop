FROM golang:1.19.1-alpine3.16 AS builder
WORKDIR /app
COPY . .
RUN go build -o main cmd/main.go

FROM alpine:3.16

WORKDIR /app
RUN mkdir media


COPY --from=builder /app/main .
COPY templates ./templates

EXPOSE 9090

CMD ["/app/main"]