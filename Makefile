DOCKER_COMPOSE = docker-compose
SERVICE_NAME = tokens-dev

BUILD_MSG = Building the $(SERVICE_NAME) service...
START_MSG = Starting the $(SERVICE_NAME) service...
STOP_MSG = Stopping the $(SERVICE_NAME) service...
REMOVE_MSG = Removing the $(SERVICE_NAME) service...
DB_START_MSG = Starting the PostgreSQL database...
DB_STOP_MSG = Stopping the PostgreSQL database...
DB_REMOVE_MSG = Removing the PostgreSQL database...

.PHONY: all build up down restart logs db-start db-stop db-remove

build:
	@echo "$(BUILD_MSG)"
	@docker compose -f ./docker-compose.yml down tokens-dev
	@docker compose -f ./docker-compose.yml build tokens-dev
	@docker compose -f ./docker-compose.yml up tokens-dev -d


up: build
	@echo "$(START_MSG)"
	@$(DOCKER_COMPOSE) up -d

generate:
	@echo "$(GENERATE_CODE_MSG)"
	@swagger generate model -f ./internal/api/api.swagger.yaml -t ./internal/api -m definition
	@swagger generate client -f ./internal/api/api.swagger.yaml -t ./internal/api  --skip-tag-packages --skip-models --existing-models=github.com/Maksim646/tokens/internal/api/definition -P models.Principal
	@swagger generate server -f ./internal/api/api.swagger.yaml -t ./internal/api/server --exclude-main --skip-tag-packages --skip-models --api-package=api --existing-models=github.com/Maksim646/tokens/internal/api/definition -P models.Principal

init_test_db:
	@echo "$(INIT_TEST_DB)"
	@docker exec -it $$(docker ps -aqf name=tokens-postgres-1) psql -U postgres -c 'CREATE DATABASE tokens_db_test'

test:
	@./test_wrapper.sh | tee test_output.log