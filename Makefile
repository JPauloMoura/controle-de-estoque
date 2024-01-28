include .env
env:
    export CONNECTION=$(POSTGRESQL_URL)

run:
	@echo "==> running infrastructure with docker"
	@docker-compose up

run-webserver:
	@echo "==> running webserver..."
	@go run ./cmd/webserver/main.go

run-api:
	@echo "==> running api..."
	@go run ./cmd/rest/main.go

kill-containers:
	@docker stop $$(docker ps -aq) && docker rm $$(docker ps -aq)

create-migration:
	@migrate create -ext sql -dir data/migrations -seq create_products_table

migrations-up: env
	@migrate -path data/migrations -database $(CONNECTION) -verbose up

migrations-down: env
	@migrate -path data/migrations -database $(CONNECTION) -verbose down

