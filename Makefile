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
	git tag $(word 2,$(MAKECMDGOALS))
	git tag api/$(word 2,$(MAKECMDGOALS))
	git tag api-server/$(word 2,$(MAKECMDGOALS))
	git tag common/$(word 2,$(MAKECMDGOALS))
	git tag gorm-gen-proto/$(word 2,$(MAKECMDGOALS))
	git tag import/$(word 2,$(MAKECMDGOALS))
	git tag tools/db/$(word 2,$(MAKECMDGOALS))
	git tag frontend/$(word 2,$(MAKECMDGOALS))

tidy:
	find . -type f -name "go.mod" -exec dirname {} \; | xargs -I {} sh -c 'echo "Running go mod tidy in {}"; cd {} && go get -u && go mod tidy'

db_scheme:
	docker compose --profile db_tools_scheme run --rm db_tools_scheme

db_cleanup:
	docker compose --profile db_tools_cleanup run --rm db_tools_cleanup

import:
	docker compose --profile import run --rm import

.PHONY: run lint test_race test version tidy db_scheme db_cleanup import
