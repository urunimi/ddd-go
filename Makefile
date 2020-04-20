.PHONY: test clean

default: run

help:
	@echo 'Management commands:'
	@echo
	@echo 'Usage:'
	@echo '    make migrate-alpha   Migrate development db.'
	@echo '    make migrate-prod    Migrate production db.'
	@echo

run:
	go run cmd/gorest/main.go

migrate-alpha:
	DATA_SOURCE_NAME=$$DATA_SOURCE_NAME go run ./cmd/agdbmanager

migrate-prod:
	DATA_SOURCE_NAME=$$DATA_SOURCE_NAME go run ./cmd/agdbmanager

clean:
	go clean && go mod tidy

test-build:
	docker-compose -f docker-compose.test.yml build

test: test-migrate
	docker-compose -f docker-compose.test.yml run test