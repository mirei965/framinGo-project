BINARY_NAME=framinGo

build:
	@go mod vendor
	@echo "Building FraminGo..."
	@go build -o tmp/${BINARY_NAME} .
	@echo "FraminGo built!"

run: build
	@echo "Starting FraminGo..."
	@./tmp/${BINARY_NAME} &
	@echo "FraminGo started!"

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
	@echo "Stopping FraminGo..."
	@-pkill -SIGTERM -f "./tmp/${BINARY_NAME}"
	@echo "Stopped FraminGo!"

restart: stop start