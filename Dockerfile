FROM golang:1.23-alpine

WORKDIR /app

COPY . .

RUN go build -o go-bot

CMD ["./go-bot"]
