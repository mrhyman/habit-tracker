# VARS
MAIN=$(PWD)/cmd/*.go
BIN_PATH=$(PWD)/bin
BINARY_NAME=bot-service
TMP=$(BIN_PATH)/tmp

# COMMANDS
.PHONY: migrate
migrate:
	cd migrations && tern migrate

.PHONY: vendor
vendor:
	go mod vendor

.PHONY: build
build:
	mkdir -p $(BIN_PATH)
	go build -mod=vendor -o=$(BIN_PATH)/$(BINARY_NAME) $(MAIN)

.PHONY: run
run:
	$(BIN_PATH)/$(BINARY_NAME)

.PHONY: up
up:
	docker compose -f docker-compose.yml --project-name bot up -d --build

.PHONY: infra
infra:
	docker compose -f docker-compose.yml --project-name bot up -d postgres pgbouncer kafka kafka-ui --build

.PHONY: down
down:
	docker compose -f docker-compose.yml --project-name bot stop

.PHONY: restart
restart:
	make down
	make up

## Команда удаляет временные файлы, кэш и т.д.
.PHONY: clean
clean:
	rm -fr $(BIN_PATH)
	rm -fr $(TMP)
	rm -fr $(PWD)/vendor
	rm -fr $(PROTO_OUT)

.PHONY: generate
generate:
	go generate ./...

.PHONY: test
test:
	go test -cover -v ./...

.PHONY: docs
docs:
	swag init -g cmd/main.go
	swag fmt
