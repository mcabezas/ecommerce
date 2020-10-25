# cross parameters
SHELL:=/bin/bash -O extglob
BINARY=api
VERSION=0.1.0

LDFLAGS=-ldflags "-X main.Version=${VERSION}"

# Build step, generates the binary.
build: clean
	@go build ${LDFLAGS} -o ${BINARY} cmd/web/rest/*.go

clean: ## Clean the project, set it up for a new build
	@rm -rf api

# Web is a mask to run the web interface, in our case the main function will start the http server.
web:
	@clear
	@go run cmd/web/main/!(*_test).go

# Run the test for all the directories.
test:
	@clear
	@go test -v ./...

