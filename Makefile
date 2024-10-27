include .env
.PHONY: docker-build
docker-build:
	docker build -t web-employee:local .
	docker build --tag migrator -f migration.Dockerfile .
local-migration-up:
	goose -dir ${MIGRATION_DIR} postgres ${PG_DSN} up -v
local-migration-down:
	goose -dir ${MIGRATION_DIR} postgres ${PG_DSN} down -v