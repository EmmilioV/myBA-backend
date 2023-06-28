# syntax=docker/dockerfile:1

FROM golang:1.19

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /bin/api ./cmd/api

ENV PORT=8080 \
    DATABASE_HOST=host \
    DATABASE_PORT=5432 \
    DATABASE_USER=user \
    DATABASE_PASSWORD=password \
    DATABASE_NAME=myba \
    DATABASE_SSLMODE=disable \
    MESSAGING_URL=amqp://guest:guest@172.17.0.2:5672

EXPOSE 8080

ENTRYPOINT ["/bin/api"]