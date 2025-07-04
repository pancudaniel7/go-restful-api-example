.PHONY: build run docker-up docker-down clean

APP_NAME=go-restful-api
DOCKER_COMPOSE_FILE=deployments/docker-compose.yml
MYSQL_CONTAINER_NAME=deployments-database-1

docker-database-up:
	@echo "Starting Docker database..."
	@docker-compose -f $(DOCKER_COMPOSE_FILE) up -d database

docker-up:
	@echo "Starting Docker containers..."
	docker-compose -f $(DOCKER_COMPOSE_FILE) up -d

docker-restart:
	@echo "Restarting Docker containers..."
	docker-compose -f $(DOCKER_COMPOSE_FILE) down && \
	docker-compose -f $(DOCKER_COMPOSE_FILE) up -d

docker-down:
	@echo "Stopping Docker containers..."
	docker-compose -f $(DOCKER_COMPOSE_FILE) down

docker-build:
	@echo "Building Docker image..."
	$(eval TAG := $(shell git describe --tags --abbrev=0 2>/dev/null || echo latest))
	docker build --pull --no-cache -t $(APP_NAME):$(TAG) .

clean:
	@echo "Cleaning up..."
	rm -f $(APP_NAME)

build:
	@echo "Building Go application..."
	go build -o $(APP_NAME) ./cmd/main.go

run: build
	@echo "Running Go application..."
	./$(APP_NAME)	
