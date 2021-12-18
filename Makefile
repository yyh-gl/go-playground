.DEFAULT_GOAL := help
.PHONY: help
help: ## helpを表示
	@echo '  see:'
	@echo '   - https://github.com/yyh-gl/go-playground'
	@echo ''
	@grep -E '^[%/a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-22s\033[0m %s\n", $$1, $$2}'

.PHONY: build
build: ## ビルド
	GOOS=linux GOARCH=amd64 go build -o ./bin/playground ./main.go

.PHONY: bench
bench: ## ベンチマーク測定
	docker-compose exec playground go test -bench . -benchmem

.PHONY: login
login: ## playgroundイメージのbashにアクセス
	docker-compose exec playground /bin/bash
