#!/bin/bash
set -euo pipefail

export PORT=19004
export ENV=test
export DB_NAME=${ENV}
export DB_USER=${ENV}
export DB_PASS=${ENV}
export DB_PORT=54321

go run ./app/infra/ent/migration/main.go

go run ./app/main.go > /dev/null 2>&1 &

# TODO remove sleep
sleep 1

go test ./app/domain/...

go test ./app/xxx/...

go test ./test/...

PID=`lsof -P -i:${PORT} | tail -n 1 | awk '{print $2}'`; kill -9 ${PID}
