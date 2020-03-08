SUBPACKAGES := $(shell go list ./...)

.DEFAULT_GOAL := help
DB_USER := $(shell cat .env | grep DB_USER | cut -d "=" -f2)
DB_PASSWORD := $(shell cat .env | grep DB_PASSWORD | cut -d "=" -f2)

.PHONY: init
init: ## Initialise application
	docker network create external_network || true
	docker-compose up -d
	sleep 15
	make install-fixtures

.PHONY: start
start: ## Start application
	docker network create external_network || true
	docker-compose up -d

.PHONY: stop
stop: ## Stop application
	docker-compose down

.PHONY: install-fixtures
install-fixtures: ## Install fixtures
	docker-compose exec -T mysql mysql -u $(DB_USER) -p$(DB_PASSWORD) < schema.sql

.PHONY: help
help: ## Help
	@grep -E '^[0-9a-zA-Z_/()$$-]+:.*?## .*$$' $(lastword $(MAKEFILE_LIST)) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

