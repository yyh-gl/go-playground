FROM golang:1.14.3 AS build
WORKDIR /go/src/playground

ENV TZ="Asia/Tokyo"
ENV GO111MODULE=on
COPY . .
RUN go get -u github.com/cosmtrek/air

FROM alpine:latest AS app
COPY --from=build /go/src/playground/cmd/playground/bin/playground /usr/local/bin/playground
EXPOSE 8080
ENTRYPOINT ["playground"]
