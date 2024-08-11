.PHONY:install
install:
	@echo "Installing backend dependencies..."
	@go get ./...
	@go mod tidy
	@go mod download
	@echo "Backend dependencies installed."

.PHONY:build-prod
build-prod:
	@echo "Building backend..."
	@go build -tags prod -o bin/momentum server/main.go

.PHONY:build-dev
build-dev:
	@echo "Building backend..."
	@go build -tags dev -o bin/momentum server/main.go

.PHONY:run
run:
	@echo "Running backend..."
	@./bin/momentum
