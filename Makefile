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
		echo "E.g.: make version v1.24.0"; \
		exit 1; \
	fi
	git tag $(word 2,$(MAKECMDGOALS))
	git tag api/$(word 2,$(MAKECMDGOALS))
	git tag common/$(word 2,$(MAKECMDGOALS))
	git tag db/$(word 2,$(MAKECMDGOALS))
	git tag import/$(word 2,$(MAKECMDGOALS))
	git tag import-cat/$(word 2,$(MAKECMDGOALS))
	git tag server/$(word 2,$(MAKECMDGOALS))
	git tag tools/db/$(word 2,$(MAKECMDGOALS))

tidy:
	find . -type f -name "go.mod" -exec dirname {} \; | xargs -I {} sh -c 'echo "Running go mod tidy in {}"; cd {} && go get -u ./... && go mod tidy'

db_backup:
	docker exec -t cheque_postgres mkdir -p /backup
	docker exec -t cheque_postgres pg_dump -U user -d cheque_db -F c -f /backup/cheque_db.dump
	docker cp cheque_postgres:/backup/cheque_db.dump ./var/backup/cheque_db.dump
	#tar -czvf ./var/backup/cheque_db_$(shell date +"%Y%m%d-%H%M%S").tar.gz -C ./var/backup cheque_db.dump
	#tar -tzf ./var/backup/cheque_db_*.tar.gz
	zip ./var/backup/cheque_db_$(shell date +"%Y%m%d-%H%M%S").zip ./var/backup/cheque_db.dump
	unzip -l ./var/backup/cheque_db_*.zip
	rm ./var/backup/cheque_db.dump

db_restore_from_backup:
	docker cp ./var/backup/cheque_db.dump cheque_postgres:/backup/cheque_db.dump
	docker exec -t cheque_postgres pg_restore -U user -d cheque_db --clean --if-exists /backup/cheque_db.dump

db_scheme:
	#docker exec -it cheque_postgres psql -U user -d postgres -c "CREATE DATABASE cheque_db;"
	docker compose --profile db_tools_scheme run --rm db_tools_scheme

db_cleanup:
	docker compose --profile db_tools_cleanup run --rm db_tools_cleanup

#db_drop_all_the_whole_database:
	#docker exec -it cheque_postgres psql -U user -d postgres -c "SELECT pg_terminate_backend(pg_stat_activity.pid) FROM pg_stat_activity WHERE datname = 'cheque_db';"
	#docker exec -it cheque_postgres psql -U user -d postgres -c "DROP DATABASE cheque_db;"

import:
	docker compose --profile import run --rm import

import_cat:
	docker compose --profile import_cat run --rm import_cat

.PHONY: run lint test_race test version tidy db_backup db_retsore_from_backup db_scheme db_cleanup import import_cat
