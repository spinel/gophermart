#!/bin/bash
export HTTP_ADDR=localhost:8092
export URI_SCHEME=http

# Postgres settings
export PG_URL=postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable
export PG_MIGRATIONS_PATH=file://../store/bun/migrations
export TRANSACTION_SUPPORT=true

# Logger settings
export LOG_LEVEL=debug
