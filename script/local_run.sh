#!/bin/bash
set -euo pipefail

export APP_NAME=xxx
export ENV=local
export DB_HOST=localhost
export DB_NAME=${APP_NAME}
export DB_USER=${APP_NAME}
export DB_PASS=${APP_NAME}
export DB_PORT=5432

go run ./app/main.go
