-include .env
export PATH := $(shell go env GOPATH)/bin:$(PATH)
export

# ─── App ─────────────────────────────────────────────────────────────────────

build:
	go build -o bin/api cmd/api/main.go
	go build -o bin/worker cmd/worker/main.go
	go build -o bin/migrate cmd/migrate/main.go

run:
	go run cmd/api/main.go

run-worker:
	go run cmd/worker/main.go

# ─── Test ─────────────────────────────────────────────────────────────────────

test:
	go test -v ./...

test-unit:
	go test -v $$(go list ./... | grep -v /test/integration)

test-integration:
	go test -v -tags=integration ./test/integration/...

tidy:
	go mod tidy

# ─── Migrations ──────────────────────────────────────────────────────────────
# DATABASE_URL can be overridden on the command line:
#   make migrate-up DATABASE_URL=postgres://...

DATABASE_URL ?= postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)

# Apply all pending migrations
migrate-up:
	go run ./cmd/migrate -action up -database "$(DATABASE_URL)"

# Roll back the last applied migration
migrate-down:
	go run ./cmd/migrate -action down -database "$(DATABASE_URL)"

# Roll back and reapply the last migration
migrate-redo:
	go run ./cmd/migrate -action redo -database "$(DATABASE_URL)"

# Show migration status
migrate-status:
	go run ./cmd/migrate -action status -database "$(DATABASE_URL)"

# Create a new migration (requires goose CLI)
# Usage: make migrate-create name=create_something_table
migrate-create:
	@test -n "$(name)" || (echo "ERROR: name is required. Usage: make migrate-create name=create_xxx_table" && exit 1)
	goose -dir migrations create "$(name)" sql

# Run Seeder
seeder-run:
	go run ./cmd/seeder/main.go
	
.PHONY: run build test test-unit test-integration tidy migrate-up migrate-down migrate-redo migrate-status migrate-create
