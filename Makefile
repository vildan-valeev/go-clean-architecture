.PHONY: help
help:
	@echo "-Комп перезагружал?\n-Да!\n-Клаву протирал?\n-Да!\n-Тогда не знаю в чем проблема..."

.PHONY: up
up:
	docker compose up --build

.PHONY: up_local_infra
up_local_infra:
	docker compose -f docker-compose.dev.yml up --build --remove-orphans

.PHONY: down_local_infra
down_local_infra:
	docker compose -f docker-compose.dev.yml down

.PHONY: migrate
migrate:
	cd migrations && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./migration .
	cd migrations && ./migration

.PHONY: up_local
up_local:
	cd app && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./build/main ./cmd/app/main.go
	cd app && ./build/main

.PHONY: test
test:
	cd app && go test ./...

.PHONY: test_race
test_race:
	cd app && go test -race -short ./...

.PHONY: generate
generate:
	cd app && go generate ./...

.PHONY: docs
docs: ### swag init
	cd app && swag init -g internal/transport/http/v1/api.go --output docs

