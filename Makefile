build:
	go build -o ./bin/app ./app/main.go
lint:
	go vet ./... && staticcheck ./...
format:
	go fmt ./...	
.PHONY: ent
ent:
	go generate ./ent
buflint:
	(cd proto && buf lint)
.PHONY: proto
proto:
	protoc ./proto/**/**/*.proto --go_out=./proto/ --go-grpc_out=./proto/
.PHONY: test
test:
	./script/local_test.sh
.PHONY: mock
mock:
	rm -rf ./mock/* && \
	go generate ./app/domain/xxx
local_init:
	docker-compose up -d
local_run:
	./script/local_run.sh
local_psql:
	docker-compose exec dev_db psql -U xxx
local_migrate:
	./script/local_migrate.sh
