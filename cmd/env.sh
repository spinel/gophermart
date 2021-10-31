#!/bin/bash
export HTTP_ADDR=localhost:8093
export URI_SCHEME=http

# Postgres settings
export PG_URL=postgres://postgres:postgres@localhost:5439/postgres?sslmode=disable
export PG_MIGRATIONS_PATH=file://../store/bun/migrations
export TRANSACTION_SUPPORT=true

#Redis
export REDIS_URL=redis://localhost

# Logger settings
export LOG_LEVEL=debug
