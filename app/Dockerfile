FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o mainApp ./app/cmd/main.go

CMD [ "./mainApp" ]