#!/bin/sh

set -e

echo "Running DB migrations"
/streamfair_token_svc/migrate -path /streamfair_token_svc/migration -database "$DB_SOURCE_TOKEN_SERVICE" -verbose up

echo "Starting Token Service"
exec "$@"