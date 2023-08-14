# Build the docker container
build:
	go mod tidy && go mod vendor && docker-compose up --build

# Stop the docker container
down:
	docker-compose down --remove-orphans --volumes
	rm -rf tmp

# Run the Go server
run:
	air

# Clean compiled files
clean:
	rm -rf tmp

# Install dependencies
install:
	go mod tidy && go mod vendor

# Generate Go functions from DB queries 
generate:
	sqlc generate

# Run DB migrations 
migrate:
	goose postgres postgres://postgres:@localhost:5432/crowdQA up

# Default target (this will be executed if you just run "make" without a command)
default: install
