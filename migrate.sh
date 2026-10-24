#!/bin/bash

set -e

# If a single argument is passed and it's a valid goose command, treat it as COMMAND
if [ "$#" -eq 1 ] && [[ "$1" =~ ^(up|down|status|redo|reset|version|up-by-one)$ ]]; then
    DOMAIN="all"
    COMMAND="$1"
else
    DOMAIN=${1:-all}
    COMMAND=${2:-up}
fi

DB_URL="postgresql://postgres:secret@localhost:5431/postgres"
MIGRATION_DIR="migration"

echo "Running goose migrations ($COMMAND) for domain: $DOMAIN..."
goose -dir "$MIGRATION_DIR" postgres "$DB_URL" "$COMMAND"