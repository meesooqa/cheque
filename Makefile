SHELL := /bin/bash

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
	find . -type f -name "go.mod" -exec dirname {} \; | xargs -I {} sh -c 'echo "Running go mod tidy in {}"; cd {} && go get -u ./... && go mod tidy'

db_backup:
	docker exec -t cheque04_postgres mkdir -p /backup
	docker exec -t cheque04_postgres pg_dump -U user -d receipts_db -F c -f /backup/receipts_db.dump
	docker cp cheque04_postgres:/backup/receipts_db.dump ./var/backup/receipts_db.dump
	#tar -czvf ./var/backup/receipts_db_$(shell date +"%Y%m%d-%H%M%S").tar.gz -C ./var/backup receipts_db.dump
	#tar -tzf ./var/backup/receipts_db_*.tar.gz
	zip ./var/backup/receipts_db_$(shell date +"%Y%m%d-%H%M%S").zip ./var/backup/receipts_db.dump
	unzip -l ./var/backup/receipts_db_*.zip
	rm ./var/backup/receipts_db.dump

db_restore_from_backup:
	docker cp ./var/backup/receipts_db.dump cheque04_postgres:/backup/receipts_db.dump
	docker exec -t cheque04_postgres pg_restore -U user -d receipts_db --clean --if-exists /backup/receipts_db.dump

db_scheme:
	#docker exec -it cheque04_postgres psql -U user -d postgres -c "CREATE DATABASE receipts_db;"
	docker compose --profile db_tools_scheme run --rm db_tools_scheme

db_cleanup:
	docker compose --profile db_tools_cleanup run --rm db_tools_cleanup

#db_drop_all_the_whole_database:
	#docker exec -it cheque04_postgres psql -U user -d postgres -c "SELECT pg_terminate_backend(pg_stat_activity.pid) FROM pg_stat_activity WHERE datname = 'receipts_db';"
	#docker exec -it cheque04_postgres psql -U user -d postgres -c "DROP DATABASE receipts_db;"

import:
	docker compose --profile import run --rm import

.PHONY: run lint test_race test version tidy db_backup db_retsore_from_backup db_scheme db_cleanup import
