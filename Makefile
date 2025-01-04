include .env

.DEFAULT_GOAL := help
# Colours used in help
GREEN    := $(shell tput -Txterm setaf 2)
WHITE    := $(shell tput -Txterm setaf 7)
YELLOW   := $(shell tput -Txterm setaf 3)
RESET    := $(shell tput -Txterm sgr0)

HELP_FUN = %help; \
	while(<>) { push @{$$help{$$2 // 'Misc'}}, [$$1, $$3] \
	if /^([a-zA-Z\-]+)\s*:.*\#\#(?:@([a-zA-Z\-]+))?\s(.*)$$/ }; \
	for (sort keys %help) { \
	print "${WHITE}$$_${RESET}\n"; \
	for (@{$$help{$$_}}) { \
	$$sep = "." x (32 - length $$_->[0]); \
	print "  ${YELLOW}$$_->[0]${RESET} $$sep ${GREEN}$$_->[1]${RESET}\n"; \
	}; \
	print "\n"; } \
	$$sep = "." x (32 - length "help"); \
	print "${WHITE}Options${RESET}\n"; \
	print "  ${YELLOW}help${RESET} $$sep ${GREEN}Prints this help${RESET}\n";

help:
	@echo "\nUsage: make ${YELLOW}<target>${RESET}\n\nThe following targets are available:\n";
	@perl -e '$(HELP_FUN)' $(MAKEFILE_LIST)

ROOT_DIR:=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
COMPOSE_FILE:=$(ROOT_DIR)/docker-compose.yml
DOCKER_COMPOSE:=docker compose -f $(COMPOSE_FILE)
MIGRATION_FOLDER:=$(ROOT_DIR)/migrations

setup: ##@Development Run and build docker containers
	$(DOCKER_COMPOSE) up -d --build

up: ##@Development Start docker containers in detached mode
	$(DOCKER_COMPOSE) up -d

down: ##@Development Stop docker containers
	$(DOCKER_COMPOSE) down

restart: ##@Development Restart docker containers
	$(DOCKER_COMPOSE) restart

migration-up: ##@Development Run up migrations
	migrate -path ${MIGRATION_FOLDER} -database ${DB_URL} up

migration-down: ##@Development Run down migrations
	migrate -path ${MIGRATION_FOLDER} -database ${DB_URL} down

logs: ##@Development Show log stream
	$(DOCKER_COMPOSE) logs -f

proxy: ##@Development Start proxy
	ngrok http 8080
