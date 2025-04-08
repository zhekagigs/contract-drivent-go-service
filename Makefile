.PHONY: generate clean build run test

# Generate server code from OpenAPI spec
generate:
	swagger generate server -f swagger.yaml -A taskmanager --target ./internal/generated

# Generate client code from OpenAPI spec
generate-client:
	swagger generate client -f ./api/swagger.yaml -A taskmanager --target ./pkg/client

# Clean generated code
clean:
	rm -rf ./internal/generated
	rm -rf ./pkg/client/generated

# Build the service
build:
	go build -o bin/taskmanager ./cmd/taskmanager-server

# Run the service
run:
	go run main.go --port=8080

# Run tests
test:
	go test ./...
