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
	cd db/migrations && goose postgres postgres://postgres:@localhost:5432/crowdQA up

# Seed users 
seed-users:
	cd cmd/seeder && go build -o ../../seeder && cd ../.. && ./seeder -seed=user && rm seeder

# Seed events 
seed-events:
	cd cmd/seeder && go build -o ../../seeder && cd ../.. && ./seeder -seed=event -user_id=$(user_id) && rm seeder

# Default target (this will be executed if you just run "make" without a command)
default: run
