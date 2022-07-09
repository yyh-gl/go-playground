FROM golang:latest
WORKDIR /go/src/github.com/yyh-gl/go-playground

ENV TZ="Asia/Tokyo"
ENV GO111MODULE=on

COPY . .

RUN go install github.com/cosmtrek/air@latest
RUN go install github.com/tenntenn/goplayground/cmd/gp@latest
