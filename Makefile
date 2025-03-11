lint:
	golangci-lint run ./...

test_race:
	go test -race -timeout=60s -count 1 ./...

test:
	go clean -testcache
	go test ./...

version:
	@if [ -z "$(word 2,$(MAKECMDGOALS))" ]; then \
		echo "E.g.: make version v1.22.33"; \
		exit 1; \
	fi
	git tag api/$(word 2,$(MAKECMDGOALS))
	git tag api-server/$(word 2,$(MAKECMDGOALS))
	git tag common/$(word 2,$(MAKECMDGOALS))
	git tag gorm-gen-proto/$(word 2,$(MAKECMDGOALS))
	git tag import/$(word 2,$(MAKECMDGOALS))
	git tag tools/db/$(word 2,$(MAKECMDGOALS))

db_init:
	go run ./tools/db gorm:migrate

db_cleanup:
	go run ./tools/db cleanup

import:
	go run ./import

.PHONY: run lint test_race test db_init db_cleanup import
