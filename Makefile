include .env

# Run tests
test:
	@echo "==> running tests"
	@go test -v ./...

# Run infrastructure with Docker Compose
run: create-network
	@echo "==> running infrastructure with docker"
	@docker-compose up

run-webserver: create-network
	@echo "==> running webserver..."
	@go run ./cmd/webserver/main.go

run-api: create-network
	@echo "==> running api..."
	@go run ./cmd/rest/main.go

run-api-with-air: create-network
	@echo "==> running api..."
	@air

# Create Docker network if it doesn't exist
create-network:
	@if ! docker network inspect $(NETWORK_NAME) >/dev/null 2>&1 ; then \
		echo "creating network $(NETWORK_NAME)..."; \
		docker network create $(NETWORK_NAME); \
	fi

kill-containers:
	docker container rm -f $$(docker container ls -qa)

# Create a new migration file
create-migration:
	@migrate create -ext sql -dir infrastructure/database/migrations -seq create_products_table

# Run migrations up
migrations-up:
	@migrate -path infrastructure/database/migrations -database $(DB_CONNECTION_STRING) -verbose up

# Rollback migrations
migrations-down:
	@migrate -path infrastructure/database/migrations -database $(DB_CONNECTION_STRING) -verbose down

