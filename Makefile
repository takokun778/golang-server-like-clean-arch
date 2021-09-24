build:
	go build -o ./bin/app ./app/main.go

buflint:
	(cd proto && buf lint)

protogen:
	protoc ./proto/**/**/*.proto --go_out=./proto/ --go-grpc_out=./proto/

dockerup:
	docker-compose up -d

devrun:
	./script/dev-run.sh

test:
	./script/test.sh

.PHONY: migratedev
devmigrate:
	./script/dev-migrate.sh

devpsql:
	docker-compose exec dev_db psql -U clean
