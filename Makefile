-include .env
export db_host := $(shell awk -F "=" '/^DB_HOST/ {print $$2}' .env)
export db_port := $(shell awk -F "=" '/^DB_PORT/ {print $$2}' .env)
export db_username := $(shell awk -F "=" '/^DB_USER/ {print $$2}' .env)
export db_password := $(shell awk -F "=" '/^DB_PASSWORD/ {print $$2}' .env)
export db_name := $(shell awk -F "=" '/^DB_NAME/ {print $$2}' .env)

db-instance-up:
	DOCKER_BUILDKIT=1 docker compose -f docker-compose-db-setup-local.yml down
	DOCKER_BUILDKIT=1 docker compose -f docker-compose-db-setup-local.yml up -d --force-recreate --remove-orphans --build

db-instance-down:
	DOCKER_BUILDKIT=1 docker compose -f docker-compose-db-setup-local.yml down

migrate:
	migrate -path migrations -database "postgres://$(db_username):$(db_password)@$(db_host):$(db_port)/$(db_name)?sslmode=disable" -verbose up

revertmigrations:
	migrate -path migrations -database "postgres://$(db_username):$(db_password)@$(db_host):$(db_port)/$(db_name)?sslmode=disable" -verbose down

generate-swagger-docs:
	swag init
