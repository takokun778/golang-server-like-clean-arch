#!/bin/sh -eu

export ENV=local
export DB_NAME=clean
export DB_USER=clean
export DB_PASS=clean
export DB_PORT=5432
export ENV=dev

go run ./app/main.go
