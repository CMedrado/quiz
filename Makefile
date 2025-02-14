# Variables
DOCKER_COMPOSE = docker-compose
DOCKER_COMPOSE_FILE = docker-compose.yml

# Build the project using Go
build:
	go build -o ./bin/main ./cmd/api/main.go

# Run the API directly with Go (no Docker)
run_api:
	go run ./cmd/api/main.go

# Run the API using Docker Compose
run_api_docker:
	docker-compose -f $(DOCKER_COMPOSE_FILE) up --build

# Run the CLI using Go
run_cli:
	go run ./cmd/cli/main.go quiz

# Run the tests
test:
	go test ./...

# Generate Swagger docs
swagger_generate:
	swag init --dir ./app --output ./docs

# Clean up binaries
clean:
	rm -rf ./bin

# Stop the Docker containers
stop:
	docker-compose down

# View logs of the services
logs:
	docker-compose logs -f

# Run everything (API and CLI)
run_all: run_api_docker run_cli
