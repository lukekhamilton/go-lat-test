FROM golang:alpine

WORKDIR /go/src/app
COPY . .

RUN go test -v
