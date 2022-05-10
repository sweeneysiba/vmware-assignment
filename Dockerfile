FROM golang:latest

LABEL maintainer="Siba Sankar <sweetsiba8@gmail.com>"

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY .env.sample .env
COPY . .

RUN go build cmd/main.go

CMD ["./main"]