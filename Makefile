lint:
	golangci-lint run ./...

test_race:
	go test -race -timeout=60s -count 1 ./...

test:
	go clean -testcache
	go test ./...

db_init:
	go run ./tools/db gorm:migrate

db_cleanup:
	go run ./tools/db cleanup

import:
	go run ./import

.PHONY: run lint test_race test db_init db_cleanup import
