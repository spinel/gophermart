#!/bin/bash
export RUN_ADDRESS=localhost:8093
export URI_SCHEME=http

# Postgres settings
export DATABASE_URI=postgres://postgres:postgres@localhost:5439/postgres?sslmode=disable
export PG_MIGRATIONS_PATH=file://store/bun/migrations
export ACCRUAL_SYSTEM_ADDRESS=http://localhost:35949

#Redis
export Redis_URL=redis://localhost

# Logger settings
export LOG_LEVEL=debug
