FROM golang:alpine

RUN apk add --no-cache git postgresql-client

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o app

CMD ["./app"]