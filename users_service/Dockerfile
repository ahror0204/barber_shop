FROM golang:1.19.1-alpine3.16 AS builder
WORKDIR /users
COPY . .

RUN apk add curl
RUN go build -o main cmd/main.go
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz


FROM alpine:3.16

WORKDIR /users

COPY --from=builder /users/main .
COPY --from=builder /app/migrate ./migrate
COPY migrations ./migrations
COPY templates ./templates

EXPOSE 8080

CMD ["/users/main"]