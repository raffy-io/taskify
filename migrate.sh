#!/bin/bash

set -e

DOMAIN=${1:-all}
COMMAND=${2:-up}

DB_URL="postgresql://postgres:secret@localhost:5431/postgres"
MIGRATION_DIR="migration"

echo "Running goose migrations ($COMMAND)..."
goose -dir "$MIGRATION_DIR" postgres "$DB_URL" "$COMMAND"