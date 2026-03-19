.PHONY: build run test lint clean migrate sqlc docker-up docker-down

# Build the API binary
build:
	go build -o bin/api ./cmd/api

# Run the API server
run:
	go run ./cmd/api

# Run tests
test:
	go test ./... -v -race

# Run tests with coverage
test-cover:
	go test ./... -v -race -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html

# Lint
lint:
	golangci-lint run ./...

# Clean build artifacts
clean:
	rm -rf bin/ coverage.out coverage.html

# Run database migrations
migrate-up:
	migrate -path migrations -database "$(DATABASE_URL)" up

migrate-down:
	migrate -path migrations -database "$(DATABASE_URL)" down 1

migrate-create:
	@read -p "Migration name: " name; \
	migrate create -ext sql -dir migrations -seq $$name

# Generate sqlc code
sqlc:
	sqlc generate

# Start infrastructure services
docker-up:
	docker compose up -d

# Stop infrastructure services
docker-down:
	docker compose down

# Start infrastructure and tail logs
docker-logs:
	docker compose up -d && docker compose logs -f

# Format code
fmt:
	go fmt ./...
	goimports -w .

# Verify dependencies
verify:
	go mod verify
	go mod tidy
