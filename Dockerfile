FROM golang:1.16.0
WORKDIR /go/src/github.com/yyh-gl/go-playground

ENV TZ="Asia/Tokyo"
ENV GO111MODULE=on

COPY . .

RUN go get -u github.com/cosmtrek/air
RUN go get github.com/tenntenn/goplayground/cmd/gp
