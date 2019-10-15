clean:
	@echo "Removing binaries"
	rm -f notes-rest-api

vendor:
	@echo "Installing all dependencies"
	go get ./...

build:
	@echo "Building notes-rest-api"
	go build -o notes-rest-api

test-unit:
	@echo "Running unit tests"
	./coverage.sh

test-all:
	@echo "Running all tests"
	./coverage.sh --intgr
