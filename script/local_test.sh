#!/bin/bash
set -euo pipefail

export ENV=test
export DB_NAME=${ENV}
export DB_USER=${ENV}
export DB_PASS=${ENV}
export DB_PORT=54321

go run ./app/infra/ent/migration/main.go

go test ./app/domain/...

go test ./app/xxx/...
