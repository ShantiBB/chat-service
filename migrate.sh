#!/bin/bash
source .env

if ! command -v goose &> /dev/null
then
    echo "goose not found. Installing..."
    go install github.com/pressly/goose/v3/cmd/goose@latest

    if ! command -v goose &> /dev/null
    then
        echo "Error: goose not installed. Check that \$GOPATH/bin is in PATH"
        exit 1
    fi

    echo "goose successfully installed!"
fi

export GOOSE_DRIVER=postgres
export GOOSE_DBSTRING="host=${POSTGRES_HOST} port=${POSTGRES_PORT} user=${POSTGRES_USER} password=${POSTGRES_PASSWORD} dbname=${POSTGRES_DB} sslmode=disable"
export GOOSE_MIGRATION_DIR=./migrations

goose "$@"