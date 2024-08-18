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

.PHONY:run
run:
	@echo "Building backend..."
	@go build -tags dev -o bin/momentum server/main.go
	@echo "Running backend..."
	@./bin/momentum

.PHONY:db-reset
db-reset:
	@echo "Resetting database..."
	@npx supabase db reset

.PHONY:db-start
db-start:
	@echo "Resetting database..."
	@npx supabase start

.PHONY:db-stop
db-stop:
	@echo "Resetting database..."
	@npx supabase stop