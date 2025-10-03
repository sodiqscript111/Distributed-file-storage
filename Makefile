build:
	# Build the Go binary into the bin/ directory
	@go build -o bin/fs

run: build
	# Run the binary
	@./bin/fs

test:
	# Run tests recursively
	@go test ./... -v
