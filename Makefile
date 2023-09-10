GOPATH:=$(shell go env GOPATH)
include env
.PHONY: env

format:
	gofmt -l -s -w .

analyzer:
	go vet ./...

test:
	go test -v ./... -cover

docker:
	docker build -t strava:latest .

db-up:
	migrate -path ./internal/db/migration -database "sqlite3://strava.db?x-migrations-table=migration_schema" -verbose up

db-down:
	migrate -path ./internal/db/migration -database "sqlite3://strava.db?x-migrations-table=migration_schema" -verbose down

generate:
	sqlc generate

run: format analyzer
	go run main.go

.PHONY: format analyzer docker db-up db-down generate run