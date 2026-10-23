#!/bin/bash

set -e

# Grab the domain name (e.g., 'home', 'about', 'users')
DOMAIN=$1
# Grab the command (defaults to 'up' if not provided)
COMMAND=${2:-up}

# Check if domain was provided
if [ -z "$DOMAIN" ]; then
  echo "Error: Domain name is required."
  echo "Usage: ./migrate.sh <domain> [up|down|status]"
  echo "Example: ./migrate.sh about up"
  exit 1
fi

DB_URL="postgresql://postgres:secret@localhost:5431/postgres"
MIGRATION_DIR="internal/$DOMAIN/migration/schema"

echo "Running goose for domain: $DOMAIN ($COMMAND)..."
goose -dir "$MIGRATION_DIR" postgres "$DB_URL" "$COMMAND"