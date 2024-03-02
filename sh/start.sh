#!/bin/sh

# This file serves as a reference for ENTRYPOINT in a Dockerfile
# as well as entrypoint with command in a docker-compose file.
set -e

echo "Starting Token Service"
exec "$@"