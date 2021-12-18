FROM golang:latest
WORKDIR /go/src/github.com/yyh-gl/go-playground

ENV TZ="Asia/Tokyo"
ENV GO111MODULE=on

COPY . .

RUN go get -u github.com/cosmtrek/air
RUN go get github.com/tenntenn/goplayground/cmd/gp
