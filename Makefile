db_init:
	go run ./tools/db gorm:migrate

db_cleanup:
	go run ./tools/db cleanup

lint:
	golangci-lint run ./...

test_race:
	go test -race -timeout=60s -count 1 ./...

test:
	go clean -testcache
	go test ./...

.PHONY: db_init db_cleanup run lint test_race test
