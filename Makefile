export CGO_ENABLED=0

.PHONY: build
build:
	GOOS=linux GOARCH=amd64 go build -o ./bin/playground ./main.go
