FROM golang:1.16-alpine

WORKDIR /app

COPY . .

RUN go mod download

CMD ["cd cmd"]

WORKDIR /app/cmd

RUN go build -o /auth

EXPOSE 9000

CMD ["/auth"]
