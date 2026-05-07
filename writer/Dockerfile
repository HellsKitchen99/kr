FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o mainWriter ./writer/cmd/main.go

CMD [ "./mainWriter" ]