.PHONY: all build env-up env-down

all: clean build env-up

build:
	@echo "Build..."
	@CGO_ENABLED=0 GOOS=linux go build -o vetch
	@docker build --tag vetch .
	@echo "Build done"

env-up:
	@echo "starting microservice..."
	docker-compose up --force-recreate -d
	@echo "Wait 2 secs"
	@sleep 2
	@echo "microservice up"

env-down:
	@echo "Stopping microservice..."
	docker-compose down 
	@echo "microservice down"

clean: env-down
	@echo "Cleaning up..."
	@docker rm -f -v `docker ps -a --no-trunc | grep "vetch" | cut -d ' ' -f 1` 2>/dev/null || true
	@docker rmi `docker images --no-trunc | grep "vetch" | cut -d ' ' -f 1` 2>/dev/null || true
	@echo "Clean up done"

	