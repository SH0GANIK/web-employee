#!/bin/bash
source .env

export MIGRATION_DSN="host=postgres port=5432 dbname=$POSTGRES_DB user=$POSTGRES_USER password=$POSTGRES_PASSWORD sslmode=disable"
export MIGRATION_DIR="./migrations"
sleep 2 && goose -dir "${MIGRATION_DIR}" postgres "${MIGRATION_DSN}" up -v