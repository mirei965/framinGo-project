BINARY_NAME=framingo

build:
	@go mod vendor
	@echo "Building Framingo..."
	@go build -o tmp/${BINARY_NAME} .
	@echo "Framingo built!"

run: build
	@echo "Starting Framingo..."
	@./tmp/${BINARY_NAME} &
	@echo "Framingo started!"

clean:
	@echo "Cleaning..."
	@go clean
	@rm tmp/${BINARY_NAME}
	@echo "Cleaned!"

start_compose:
	docker-compose up -d

stop_compose:
	docker-compose down 

test:
	@echo "Testing..."
	@go test ./...
	@echo "Done!"

start: run

stop:
	@echo "Stopping Framingo..."
	@-pkill -SIGTERM -f "./tmp/${BINARY_NAME}"
	@echo "Stopped Framingo!"

restart: stop start