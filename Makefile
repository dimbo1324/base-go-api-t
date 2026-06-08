-include .env
-include .envrc

MIGRATE ?= migrate
MIGRATIONS_PATH ?= ./cmd/migrate/migrations
DB_MIGRATOR_ADDR ?= postgres://postgres:password@localhost/appdb?sslmode=disable

.PHONY: migration migrate-up migrate-down test vet fmt fmt-check

migration:
	@$(MIGRATE) create -seq -ext sql -dir $(MIGRATIONS_PATH) $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@$(MIGRATE) -path=$(MIGRATIONS_PATH) -database=$(DB_MIGRATOR_ADDR) up

migrate-down:
	@$(MIGRATE) -path=$(MIGRATIONS_PATH) -database=$(DB_MIGRATOR_ADDR) down $(filter-out $@,$(MAKECMDGOALS))

test:
	@go test ./...

vet:
	@go vet ./...

fmt:
	@gofmt -w $$(find . -name '*.go' -not -path './vendor/*')

fmt-check:
	@test -z "$$(gofmt -l $$(find . -name '*.go' -not -path './vendor/*'))"
