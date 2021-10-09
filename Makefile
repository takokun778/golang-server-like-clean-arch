build:
	go build -o ./bin/app ./app/main.go
.PHONY: ent
ent:
	go generate ./ent
buflint:
	(cd proto && buf lint)
protogen:
	protoc ./proto/**/**/*.proto --go_out=./proto/ --go-grpc_out=./proto/
local_dev_init:
	docker-compose up -d
local_dev_run:
	./script/local_run.sh
local_dev_psql:
	docker-compose exec dev_db psql -U clean
local_migrate:
	./script/local_migrate.sh
test:
	./script/local_test.sh
.PHONY: mock
mock:
	rm -rf ./mock/* && \
	go generate ./app/domain/fuga && \
	go generate ./app/domain/hoge
